package config

import "log"

func (configStruct *ConfigStruct) AddListener(listener *Listener) {

	if listener.Name == "" {
		log.Fatal("Listener Name cannot be empty")
	}

	configStruct.Listeners = append(configStruct.Listeners, listener)

}
