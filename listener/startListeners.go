package listener

import (
	"log"
	"sync"
)

func StartListeners(serversWG *sync.WaitGroup) {
	log.Println("Starting listeners...")

	if len(Listeners) == 0 {
		log.Println("No listeners to start")
		return
	}

	for _, listener := range Listeners {
		if listener.ListenerDetails.Listening {
			serversWG.Add(1)
			listener.ServerHandler.Listen(serversWG, listener.ListenerDetails)
		}
	}
}
