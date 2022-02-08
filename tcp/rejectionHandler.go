package tcp

import "net"

func rejectionHandler(src net.Conn) {
	src.Write([]byte("Service Unavailable"))
	src.Close()
}
