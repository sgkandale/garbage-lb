package http

import (
	"context"
	"log"
)

func (server *Server) Terminate(wg *sync.WaitGroup, name string) {
	log.Printf("shutting down %s\n...", name)

	// waitgroup is already defered to done in listen

	err := server.Shutdown(context.TODO())
	if err != nil {
		panic(err)
	}
}
