package http

import (
	goHttp "net/http"

	"garbagelb/config"
)

type HTTPServer struct {
	Server   goHttp.Server
	Listener *config.Listener
}

func InitServer() HTTPServer {
	return HTTPServer{}
}

var Server = InitServer()
