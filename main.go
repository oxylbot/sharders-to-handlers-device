package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/zeromq/goczmq"
)

func getAddresses() (string, string) {
	// Get the addresses from the env, in format tcp://x.x.x.x:y
	incomingBind := os.Getenv("INCOMING_ADDRESS")
	outgoingBind := os.Getenv("OUTGOING_ADDRESS")

	if incomingBind == "" || outgoingBind == "" {
		// If we have a missing address we should exit
		log.Println("INCOMING_ADDRESS and OUTGOING_ADDRESS are not both set, please check them")
		os.Exit(1)
	}

	return incomingBind, outgoingBind
}

func getTypes() (int, int) {
	// Get the types of socket
	incomingTypeString := os.Getenv("INCOMING_TYPE")
	outgoingTypeString := os.Getenv("OUTGOING_TYPE")

	if incomingTypeString == "" || outgoingTypeString == "" {
		// If we have a missing type we should exit
		log.Println("INCOMING_TYPE and OUTGOING_TYPE are not both set, please check them")
		os.Exit(1)
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
	// -verbose enables the proxy.Verbose() logging
	verbosePtr := flag.Bool("verbose", false, "Enable verbose logging")

	// -capture tcp://*:xxxx/ sets up a push socket to capture all traffic going through the proxy on
	captureAddr := flag.String("capture", "", "Start a PUSH bound to this address")

	// Parse command line arguments
	flag.Parse()

	incomingBind, outgoingBind := getAddresses()

	incomingType, outgoingType := getTypes()

	// Construct CZMQ proxy
	proxy := goczmq.NewProxy()

	if *verbosePtr == true {
		// Set debug logging on
		proxy.Verbose()
	}

	if *captureAddr != "" {
		// Set the capture mode
		proxy.SetCapture(*captureAddr)
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
	<-signalChannel

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
