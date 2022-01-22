package listener

import (
	"log"
	"sync"
)

func StopListeners(serversWG *sync.WaitGroup) {
	log.Println("Stopping listeners...")

	for _, listener := range Listeners {
		listener.ServerHandler.Terminate(serversWG, listener.Name)
	}
}
