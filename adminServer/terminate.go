package adminServer

import (
	"context"
	"log"
	"sync"
)

func (webServer *AdminServer) Terminate(wg *sync.WaitGroup, name string) {
	log.Printf("shutting down %s\n...", name)

	// waitgroup is already defered to done in listen

	err := webServer.Shutdown(context.TODO())
	if err != nil {
		panic(err)
	}
}
