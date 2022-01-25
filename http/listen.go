package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

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

	// run health checks in loop to keep the health checker alive
	loopCtx, cancelLoopCtx := context.WithCancel(context.TODO())
	go func(listener *config.Listener, ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				PerformHealthChecks(listener)
				time.Sleep(time.Duration(listener.HealthCheckInterval) * time.Second)
			}
		}
	}(server.Listener, loopCtx)

	go func() {
		defer wg.Done()
		// defer
		if server.Listener.TLS {
			log.Println(
				server.ListenAndServeTLS(
					server.Listener.CertPath,
					server.Listener.KeyPath,
				),
			)
		} else {
			log.Println(server.ListenAndServe())
		}
		cancelLoopCtx()
		server.Terminate(wg, server.Listener)
	}()
}
