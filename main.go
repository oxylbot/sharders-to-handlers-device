package main

import (
	"github.com/zeromq/goczmq"
	"os"
	"os/signal"
	"syscall"
)

func getAddresses() (string, string) {
	shardersBind := os.Getenv("PUSH_ADDRESS")
	handlersBind := os.Getenv("PULL_ADDRESS")

	if shardersBind == "" || handlersBind == "" {
		panic("PUSH_ADDRESS and PULL_ADDRESS are not both set, please check them")
	}

	return shardersBind, handlersBind;
}

func main() {
	shardersBind, handlersBind := getAddresses();

	verboseLogging := os.Getenv("PROXY_DEBUG")

	proxy := goczmq.NewProxy()

	if verboseLogging == "1" {
		proxy.Verbose()
	}

	err := proxy.SetBackend(goczmq.Push, shardersBind)
	if err != nil {
		proxy.Destroy()
		panic(err)
	}

	err = proxy.SetFrontend(goczmq.Pull, handlersBind)
	if err != nil {
		proxy.Destroy()
		panic(err)
	}

	signalChannel := makeSignalChannel()
	<- signalChannel

	proxy.Destroy()
}

func makeSignalChannel() (chan os.Signal) {
	signalChannel := make(chan os.Signal, 1)

	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	return signalChannel
}