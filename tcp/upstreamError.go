package tcp

import "net"

func upstreamErrorHandler(src net.Conn) {
	src.Write([]byte("upstream connection error"))
	src.Close()
}
