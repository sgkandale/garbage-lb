package config

import "log"

func verify(config *ConfigStruct) *ConfigStruct {

	// Server Check
	if config.Server.Port == "" {
		config.Server.Port = "8080"
	}
	if config.Server.TLS {
		if config.Server.CertPath == "" {
			log.Fatal("TLS is enabled but no cert path is set")
		}
		if config.Server.KeyPath == "" {
			log.Fatal("TLS is enabled but no key path is set")
		}
	}

	// AdminPortal Check
	// if config.AdminPortal.Enabled {
	if config.AdminPortal.Port == "" {
		config.AdminPortal.Port = "8081"
	}
	// }

	return config
}

var Config = verify(parsedConfig)
