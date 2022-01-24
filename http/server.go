package http

import (
	"net/http"

	"garbagelb/config"
)

type HTTPServer struct {
	http.Server
	Listener *config.Listener
}

func InitServer() HTTPServer {
	return HTTPServer{}
}

var Server = InitServer()
