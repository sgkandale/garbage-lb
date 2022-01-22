package config

import (
	"log"

	"garbagelb/defaults"
)

func verify(config *ConfigStruct) *ConfigStruct {

	newConfig := &ConfigStruct{}

	// Admin Checks
	if config.Admin.Enabled {
		newConfig.Admin.Enabled = true
		if config.Admin.Port == 0 {
			newConfig.Admin.Port = defaults.AdminPort
		} else {
			newConfig.Admin.Port = config.Admin.Port
		}
	}

	// Cluster Checks
	if len(config.Clusters) == 0 && len(config.Listeners) > 0 {
		log.Println("No clusters defined for listeners.")
	} else {
		for _, eachCluster := range config.Clusters {
			newConfig.AddCluster(eachCluster)
		}
	}

	// Listener Checks
	if len(config.Listeners) == 0 {
		if !newConfig.Admin.Enabled {
			log.Println("No listeners defined.")
		}
	} else {
		for _, eachListener := range config.Listeners {
			newConfig.AddListener(eachListener)
		}
	}

	return newConfig
}

var Config = verify(parsedConfig)
