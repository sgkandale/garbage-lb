package tcp

import (
	"log"
	"sync"

	"garbagelb/config"
)

func (server *TCPServer) Terminate(wg *sync.WaitGroup, listener *config.Listener) {

	// waitgroup is already defered to done in listen

	// is server is not nil, close it
	if server.Server != nil {
		log.Printf("shutting down %s...", listener.Name)
		server.TerminateSignalled = true
		server.Server.Close()
	}
}
