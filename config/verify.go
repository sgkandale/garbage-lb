package config

import (
	"log"

	"garbagelb/defaults"
)

func verify(config *ConfigStruct) *ConfigStruct {

	// Admin Checks
	if config.Admin.Port == "" {
		config.Admin.Port = defaults.AdminPort
	}

	// Listener Checks
	if len(config.Listeners) == 0 {
		log.Println("No listeners defined.")
	}

	return config
}

var Config = verify(parsedConfig)
