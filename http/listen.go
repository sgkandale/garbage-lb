package http

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"garbagelb/config"
)

func (server *HTTPServer) Listen(wg *sync.WaitGroup, listener *config.Listener) {

	server.Addr = fmt.Sprintf(":%d", listener.Port)
	server.Handler = http.HandlerFunc(server.LBHandler)
	server.Listener = listener

	log.Println(
		fmt.Sprintf(
			"%s starting at port %d...",
			listener.Name,
			listener.Port,
		),
	)

	go func() {
		defer wg.Done()
		log.Fatal(server.ListenAndServe())
	}()
}
