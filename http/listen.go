package http

import (
	"context"
	"fmt"
	"log"
	goHttp "net/http"
	"sync"
	"time"

	"garbagelb/config"
)

func (server *HTTPServer) Listen(wg *sync.WaitGroup, listener *config.Listener) {

	server.Server.Addr = fmt.Sprintf(":%d", listener.Port)
	server.Server.Handler = goHttp.HandlerFunc(server.LBHandler)
	server.Listener = listener

	log.Println(
		fmt.Sprintf(
			"%s of type {%s} starting at port %d...",
			listener.Name,
			listener.Type,
			listener.Port,
		),
	)

	// run health checks in loop to keep the health checker alive
	loopCtx, cancelLoopCtx := context.WithCancel(context.Background())
	go func(listener *config.Listener, ctx context.Context) {
		// wait if the server shuts down due to some error
		time.Sleep(time.Second * 2)
		for {
			select {
			case <-ctx.Done():
				log.Printf(
					"health checks for listener {%s} stopped",
					listener.Name,
				)
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
				server.Server.ListenAndServeTLS(
					server.Listener.CertPath,
					server.Listener.KeyPath,
				),
			)
		} else {
			log.Println(server.Server.ListenAndServe())
		}
		cancelLoopCtx()
		server.Terminate(wg, server.Listener)
	}()
}
