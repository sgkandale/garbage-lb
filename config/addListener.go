package config

import (
	"log"
	"strings"
)

func (configStruct *ConfigStruct) AddListener(listener *Listener) {

	if listener.Name == "" {
		log.Fatal("Listener Name cannot be empty")
	}

	for _, eachListener := range configStruct.Listeners {
		// check case insensitive name
		if strings.EqualFold(eachListener.Name, listener.Name) {
			log.Fatalf("listener name {%s} is already in use", listener.Name)
		}
	}

	configStruct.Listeners = append(configStruct.Listeners, listener)

}
