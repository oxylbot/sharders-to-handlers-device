package main

import (
	"github.com/zeromq/goczmq"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func getAddresses() (string, string) {
	incomingBind := os.Getenv("INCOMING_ADDRESS")
	outgoingBind := os.Getenv("OUTGOING_ADDRESS")

	if incomingBind == "" || outgoingBind == "" {
		panic("INCOMING_ADDRESS and OUTGOING_ADDRESS are not both set, please check them")
	}

	return incomingBind, outgoingBind
}

func getTypes() (int, int) {
	incomingTypeString := os.Getenv("INCOMING_TYPE")
	outgoingTypeString := os.Getenv("OUTGOING_TYPE")

	if incomingTypeString == "" || outgoingTypeString == "" {
		panic("INCOMING_TYPE and OUTGOING_TYPE are not both set, please check them")
	}

	incomingType, err := strconv.Atoi(incomingTypeString)
	if err != nil {
		panic(err)
	}

	outgoingType, err := strconv.Atoi(outgoingTypeString)
	if err != nil {
		panic(err)
	}

	return incomingType, outgoingType
}

func main() {
	incomingBind, outgoingBind := getAddresses()
	
	incomingType, outgoingType := getTypes()

	verboseLogging := os.Getenv("PROXY_DEBUG")

	proxy := goczmq.NewProxy()

	if verboseLogging == "1" {
		proxy.Verbose()
	}

	err := proxy.SetBackend(incomingType, incomingBind)
	if err != nil {
		proxy.Destroy()
		panic(err)
	}

	err = proxy.SetFrontend(outgoingType, outgoingBind)
	if err != nil {
		proxy.Destroy()
		panic(err)
	}

	signalChannel := makeSignalChannel()
	<- signalChannel

	proxy.Destroy()
}

func makeSignalChannel() chan os.Signal {
	signalChannel := make(chan os.Signal, 1)

	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	return signalChannel
}