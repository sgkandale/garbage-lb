package listener

import (
	"garbagelb/adminServer"
	"garbagelb/config"
)

func ListListeners() int {
	totalListeners := 0

	if config.Config.Admin.Enabled {
		totalListeners++
		adminListener := Listener{
			ServerHandler: &adminServer.Server,
			ListenerDetails: &config.Listener{
				ID:        "admin",
				Name:      "admin server",
				Port:      config.Config.Admin.Port,
				Type:      "http",
				Listening: true,
			},
		}
		Listeners = append(Listeners, &adminListener)
	}

	for _, listener := range config.Config.Listeners {
		totalListeners++
		newListener := Listener{
			ServerHandler: nil,
			ListenerDetails: &config.Listener{
				Name: listener.Name,
				Port: listener.Port,
			},
		}
		Listeners = append(Listeners, &newListener)
	}

	return totalListeners
}
