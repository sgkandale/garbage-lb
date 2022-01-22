package main

import (
	"log"
	"os"
	"os/signal"
	"sync"

	"garbagelb/listener"
)

func main() {

	log.Println("Starting GarbageLB...")

	serversWG := &sync.WaitGroup{}

	registeredListeners := listener.ListListeners()
	log.Printf("total listeners : %d\n", registeredListeners)

	listener.StartListeners(serversWG)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch

	// attempt to stop all servers
	listener.StopListeners(serversWG)

	// wait for all servers to stop
	serversWG.Wait()

	log.Println("Stopping GarbageLB...")

}
