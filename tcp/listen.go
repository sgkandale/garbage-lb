package tcp

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"garbagelb/config"
)

func (server *TCPServer) Listen(wg *sync.WaitGroup, listener *config.Listener) {

	server.Listener = listener
	newServer, err := net.Listen(
		listener.Type,
		fmt.Sprintf(":%d", listener.Port),
	)
	if err != nil {
		log.Printf(
			"error starting tcp listener {%s} : %s",
			listener.Name,
			err.Error(),
		)
		wg.Done()
		return
	}
	server.Server = newServer

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

	// accept connections in seperate goroutine
	go func(listener *config.Listener, ctx context.Context) {
	connectionLoop:
		for {
			if server.TerminateSignalled {
				log.Printf(
					"%s of type {%s} shutting down...",
					listener.Name,
					listener.Type,
				)
				cancelLoopCtx()
				wg.Done()
				break connectionLoop
			} else {
				// wait for a connection
				conn, err := server.Server.Accept()
				if err != nil {
					log.Printf(
						"error accepting connection in listener {%s} : %s",
						listener.Name,
						err.Error(),
					)
					// continue with loop
					continue connectionLoop
				}

				// handle the connection in a new goroutine
				go func(newConn net.Conn) {
					// lb handler
					server.LBHandler(newConn)
					// newConn.Close()
				}(conn)
			}
		}
	}(server.Listener, loopCtx)

}
