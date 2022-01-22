package http

import (
	"context"
	"log"
	"sync"

	"garbagelb/config"
)

func (server *HTTPServer) Terminate(wg *sync.WaitGroup, listener *config.Listener) {
	log.Printf("shutting down %s\n...", listener.Name)

	// waitgroup is already defered to done in listen

	err := server.Shutdown(context.TODO())
	if err != nil {
		panic(err)
	}
}
