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
			Name:          "adminServer",
			Port:          config.Config.Admin.Port,
			ServerHandler: &adminServer.Server,
		}
		Listeners = append(Listeners, &adminListener)
	}

	for _, listener := range config.Config.Listeners {
		totalListeners++
		listener := Listener{
			Name:          listener.Name,
			Port:          listener.Port,
			ServerHandler: nil,
		}
		Listeners = append(Listeners, &listener)
	}

	return totalListeners
}
