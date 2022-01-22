package config

import "fmt"

func (configStruct *ConfigStruct) AddListener(listener Listener) error {

	if listener.ID == "" {
		return fmt.Errorf("Listener ID cannot be empty")
	}

	configStruct.Listeners = append(configStruct.Listeners, listener)

	return nil
}
