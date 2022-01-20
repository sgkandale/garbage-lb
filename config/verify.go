package config

import (
	"log"

	"simplelb/constants"
)

func verify(config *ConfigStruct) *ConfigStruct {

	// Admin Checks
	if config.Admin.Port == "" {
		config.Admin.Port = constants.AdminPort
	}

	// Listener Checks
	if len(config.Listeners) == 0 {
		log.Println("No listeners defined.")
	}

	return config
}

var Config = verify(parsedConfig)
