package config

import (
	"log"
	"strings"

	"garbagelb/defaults"
)

func (configStruct *ConfigStruct) AddListener(listener *Listener) {

	newListener := &Listener{}

	if listener.Name == "" {
		log.Fatal("Listener Name cannot be empty")
	}

	for _, eachListener := range configStruct.Listeners {
		// check case insensitive name
		if strings.EqualFold(eachListener.Name, listener.Name) {
			log.Fatalf("listener name {%s} is already in use", listener.Name)
		}
	}
	newListener.Name = listener.Name

	if listener.Port < 1 || listener.Port > 65535 {
		log.Fatalf(
			"out of range port {%d} for listener {%s}",
			listener.Port,
			listener.Name,
		)
	}
	newListener.Port = listener.Port

	// TLS checks pending

	for _, listenerType := range defaults.ListenerTypes {
		if strings.EqualFold(listenerType, listener.Type) {
			newListener.Type = listenerType
		}
	}
	if newListener.Type == "" {
		log.Fatalf(
			"unsupported listener type {%s} for listener {%s}",
			listener.Type,
			listener.Name,
		)
	}

	newListener.Listening = listener.Listening

	// filter checks
	for newFilterIndex, newFilter := range listener.Filters {
		// check name validity
		if newFilter.Name == "" {
			log.Fatalf(
				"filter name cannot be empty for listener {%s} at index {%d}",
				listener.Name,
				newFilterIndex,
			)
		}
		// check for existing name
		for existingFiltersIndex, existingFilter := range newListener.Filters {
			if strings.EqualFold(existingFilter.Name, newFilter.Name) {
				log.Fatalf(
					"filter name {%s} at index {%d} is already defined at index {%d} for listener {%s}",
					newFilter.Name,
					newFilterIndex,
					existingFiltersIndex,
					listener.Name,
				)
			}
		}
	}

	// add newListener to config
	configStruct.Listeners = append(configStruct.Listeners, newListener)

}
