package adminServer

import (
	"context"
	"log"
	"sync"
)

func (webServer *AdminServer) Terminate(wg *sync.WaitGroup) {
	log.Println("shutting down admin server...")

	// waitgroup is already defered to shutdown in listen

	err := webServer.Shutdown(context.TODO())
	if err != nil {
		log.Println("here")
		panic(err)
	}
}
