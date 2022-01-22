package listener

import (
	"log"
	"sync"
)

func StartListeners(serversWG *sync.WaitGroup) {
	log.Println("Starting listeners...")

	for _, listener := range Listeners {
		serversWG.Add(1)
		listener.ServerHandler.Listen(serversWG)
	}
}
