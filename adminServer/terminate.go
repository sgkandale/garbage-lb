package adminServer

import (
	"context"
	"log"
	"sync"

	"garbagelb/config"
)

func (webServer *AdminServer) Terminate(wg *sync.WaitGroup, listener *config.Listener) {
	log.Printf("shutting down %s", listener.Name)

	// waitgroup is already defered to done in listen

	err := webServer.Shutdown(context.TODO())
	if err != nil {
		panic(err)
	}
}
