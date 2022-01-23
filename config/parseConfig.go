package config

import (
	"flag"
	"log"
	"strings"

	"github.com/spf13/viper"
)

func parseFlags() string {
	configFilePath := flag.String("config", ".", "path to config file")

	flag.Parse()

	return *configFilePath
}

var configFilePath = parseFlags()

func parseConfig() *ConfigStruct {

	viper.SetConfigName("config")       // name of config file
	viper.SetConfigType("yaml")         // extension of config file
	viper.AddConfigPath(configFilePath) // optional directory for config file
	err := viper.ReadInConfig()         // Find and read the config file

	if err != nil {
		if strings.Contains(err.Error(), "Not Found") {
			log.Println("Config file not found. Using default values.")
			return &ConfigStruct{
				Admin: &Admin{
					Enabled: true,
				},
			}
		}
		log.Fatal("error parsing config file : ", err)
	}

	log.Println("parsing config file")

	var readConfig ConfigStruct

	err = viper.Unmarshal(&readConfig)
	if err != nil {
		log.Fatal("error unmarshing config file in struct : ", err)
	}

	log.Println("config file parsed successfully")

	return &readConfig
}

var parsedConfig = parseConfig()
