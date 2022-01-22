package config

import "fmt"

func (configStruct *ConfigStruct) AddListener(listener Listener) error {

	if listener.Name == "" {
		return fmt.Errorf("Listener Name cannot be empty")
	}

	configStruct.Listeners = append(configStruct.Listeners, listener)

	return nil
}
