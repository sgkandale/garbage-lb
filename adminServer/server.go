package adminServer

import (
	"net/http"
)

type AdminServer struct {
	http.Server
}

func InitServer() AdminServer {
	return AdminServer{}
}

var Server = InitServer()
