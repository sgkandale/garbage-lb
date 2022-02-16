package tcp

import (
	"fmt"
	"io"
	"log"
	"net"

	"garbagelb/config"
)

func forwardRequest(src net.Conn, endpoint *config.Endpoint) {

	// increment active connections counter
	endpoint.IncrementActiveConnectionsCounter()

	// forward request

	dst, err := net.Dial(
		endpoint.Protocol,
		fmt.Sprintf("%s:%d", endpoint.Address, endpoint.Port),
	)
	if err != nil {
		log.Printf(
			"Dial Error : %s",
			err.Error(),
		)
	}

	done := make(chan struct{})

	go func() {
		defer src.Close()
		if dst != nil {
			defer dst.Close()
			io.Copy(dst, src)
		}
		done <- struct{}{}
	}()

	go func() {
		defer src.Close()
		if dst != nil {
			defer dst.Close()
			io.Copy(src, dst)
		}
		done <- struct{}{}
	}()

	<-done
	<-done

	// decrement active connections counter
	endpoint.DecrementActiveConnectionsCounter()
}
