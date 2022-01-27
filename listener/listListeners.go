package listener

import (
	"fmt"
	"log"

	"garbagelb/adminServer"
	"garbagelb/config"
	"garbagelb/http"
	"garbagelb/tcp"
)

func ListListeners() int {
	totalListeners := 0

	if config.Config.Admin != nil {
		if config.Config.Admin.Enabled {
			totalListeners++
			adminListener := Listener{
				ServerHandler: &adminServer.Server,
				ListenerDetails: &config.Listener{
					Name:      "admin server",
					Port:      config.Config.Admin.Port,
					Type:      "http",
					Listening: true,
				},
			}
			Listeners = append(Listeners, &adminListener)
		}
	}

	for _, listener := range config.Config.Listeners {
		totalListeners++
		newListener := Listener{
			ListenerDetails: listener,
		}
		switch listener.Type {
		case "http":
			newListener.ServerHandler = &http.Server
		case "tcp", "tcp4", "tcp6":
			newListener.ServerHandler = &tcp.Server
		default:
			log.Fatal(
				fmt.Sprintf(
					"listener type {%s} is not supported",
					listener.Type,
				),
			)
		}
		Listeners = append(Listeners, &newListener)
	}

	return totalListeners
}
