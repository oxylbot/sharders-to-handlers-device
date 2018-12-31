package main

import (
	"github.com/zeromq/goczmq"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func getAddresses() (string, string) {
	// Get the addresses from the env, in format tcp://x.x.x.x:y
	incomingBind := os.Getenv("INCOMING_ADDRESS")
	outgoingBind := os.Getenv("OUTGOING_ADDRESS")

	if incomingBind == "" || outgoingBind == "" {
		panic("INCOMING_ADDRESS and OUTGOING_ADDRESS are not both set, please check them")
	}

	return incomingBind, outgoingBind
}

func getTypes() (int, int) {
	// Get the types of socket
	incomingTypeString := os.Getenv("INCOMING_TYPE")
	outgoingTypeString := os.Getenv("OUTGOING_TYPE")

	if incomingTypeString == "" || outgoingTypeString == "" {
		panic("INCOMING_TYPE and OUTGOING_TYPE are not both set, please check them")
	}

	// Convert to int in preparation of sending to CZMQ
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

	// Construct CZMQ proxy
	proxy := goczmq.NewProxy()

	if verboseLogging == "1" {
		// Set debug logging on
		proxy.Verbose()
	}

	// incomingType is as documented in README.md
	err := proxy.SetBackend(incomingType, incomingBind)
	if err != nil {
		proxy.Destroy()
		panic(err)
	}

	// outgoingType is as documented in README.md
	err = proxy.SetFrontend(outgoingType, outgoingBind)
	if err != nil {
		proxy.Destroy()
		panic(err)
	}

	// Block until one of the signals listened for is received
	signalChannel := makeSignalChannel()
	<- signalChannel

	// Close proxy and exit
	proxy.Destroy()
}

func makeSignalChannel() chan os.Signal {
	// Make a channel for signals
	signalChannel := make(chan os.Signal, 1)

	// Notify on SIGINT, SIGTERM and interrupts
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	return signalChannel
}