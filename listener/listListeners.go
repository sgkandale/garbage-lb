package listener

import (
	"fmt"

	"garbagelb/adminServer"
	"garbagelb/config"
	"garbagelb/http"
)

func ListListeners() int {
	totalListeners := 0

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

	for _, listener := range config.Config.Listeners {
		totalListeners++
		newListener := Listener{
			ServerHandler: &http.Server,
			ListenerDetails: &config.Listener{
				Name:          listener.Name,
				Port:          listener.Port,
				Type:          listener.Type,
				Listening:     listener.Listening,
				TargetCluster: listener.TargetCluster,
			},
		}
		for _, cluster := range config.Config.Clusters {
			if cluster.Name == listener.TargetCluster {
				newListener.ListenerDetails.TargetClusterDetails = cluster
			}
		}
		if newListener.ListenerDetails.TargetClusterDetails == nil {
			panic(
				fmt.Sprintf(
					"cluster : %s not found for listener : %s",
					listener.TargetCluster,
					listener.Name,
				),
			)
		}
		Listeners = append(Listeners, &newListener)
	}

	return totalListeners
}
