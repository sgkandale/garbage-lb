package cmd

import (
	"log"
	"os"
	"os/signal"
	"sync"

	"garbagelb/config"
	"garbagelb/listener"
)

func operator() {

	log.Println("Starting GarbageLB...")

	config.InitializeConfig()

	listenersWG := &sync.WaitGroup{}

	registeredListeners := listener.ListListeners()
	log.Printf("total listeners : %d\n", registeredListeners)

	listener.StartListeners(listenersWG)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch

	// attempt to stop all listeners
	listener.StopListeners(listenersWG)

	// wait for all listeners to stop
	listenersWG.Wait()

	log.Println("Stopping GarbageLB...")

}
