package adminServer

import (
	"garbagelb/config"
	"net/http"
)

type AdminServer struct {
	http.Server
}

func InitServer() AdminServer {
	server := AdminServer{}
	server.Addr = ":" + config.Config.Admin.Port
	return server
}

var Server = InitServer()
