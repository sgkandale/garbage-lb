package http

import "net/http"

type HTTPServer struct {
	http.Server
}

func InitServer() HTTPServer {
	return HTTPServer{}
}

var Server = InitServer()
