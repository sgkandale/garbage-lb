package listener

import (
	"garbagelb/adminServer"
	"garbagelb/config"
	"garbagelb/http"
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
			ServerHandler:   &http.Server,
			ListenerDetails: listener,
		}
		Listeners = append(Listeners, &newListener)
	}

	return totalListeners
}
