package listener

import (
	"log"
	"sync"
)

func StopListeners(serversWG *sync.WaitGroup) {
	if len(Listeners) > 0 {
		log.Println("Stopping listeners...")

		for _, listener := range Listeners {
			listener.ServerHandler.Terminate(serversWG, listener.ListenerDetails)
		}
	}
}
