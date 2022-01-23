package config

import (
	"log"

	"garbagelb/defaults"
)

func verify(config *ConfigStruct) *ConfigStruct {

	newConfig := &ConfigStruct{}

	// Admin Checks
	if config.Admin.Enabled {
		newAdmin := &Admin{}
		newAdmin.Enabled = true
		if config.Admin.Port == 0 {
			newAdmin.Port = defaults.AdminPort
		} else {
			newAdmin.Port = config.Admin.Port
		}
		newConfig.Admin = newAdmin
	}

	// Cluster Checks
	if len(config.Clusters) == 0 && len(config.Listeners) > 0 {
		log.Println("No clusters defined for listeners.")
	} else {
		for _, eachCluster := range config.Clusters {
			err := newConfig.AddCluster(eachCluster)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// Listener Checks
	if len(config.Listeners) == 0 {
		if !newConfig.Admin.Enabled {
			log.Println("No listeners defined.")
		}
	} else {
		for _, eachListener := range config.Listeners {
			err := newConfig.AddListener(eachListener)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return newConfig
}

var Config = verify(parsedConfig)
