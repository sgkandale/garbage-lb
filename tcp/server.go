package tcp

import (
	"net"

	"garbagelb/config"
)

type TCPServer struct {
	Server             net.Listener
	Listener           *config.Listener
	TerminateSignalled bool
}

func InitServer() TCPServer {
	return TCPServer{}
}

var Server = InitServer()
