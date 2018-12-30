package main

import (
	"github.com/zeromq/goczmq"
	"os"
)

func main() {
	shardersBind := os.Getenv("PUSH_ADDRESS")
	handlersBind := os.Getenv("PULL_ADDRESS")

	verboseLogging := os.Getenv("PROXY_DEBUG")

	proxy := goczmq.NewProxy()

	if verboseLogging == "1" {
		proxy.Verbose()
	}

	defer proxy.Destroy()

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
}