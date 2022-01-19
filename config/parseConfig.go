package config

import (
	"flag"
	"log"

	"github.com/spf13/viper"
)

func parseFlags() string {
	configFilePath := flag.String("config", ".", "path to config file")

	flag.Parse()

	return *configFilePath
}

var configFilePath = parseFlags()

func parseConfig() *Config {

	log.Println("parsing config file")

	viper.SetConfigName("config")       // name of config file
	viper.SetConfigType("yaml")         // extension of config file
	viper.AddConfigPath(configFilePath) // optional directory for config file
	err := viper.ReadInConfig()         // Find and read the config file

	if err != nil {
		log.Fatal("error parsing config file : ", err)
	}

	var readConfig Config

	err = viper.Unmarshal(&readConfig)
	if err != nil {
		log.Fatal("error unmarshing config file in struct : ", err)
	}

	log.Println("config file parsed successfully")

	return readConfig
}

var parsedConfig = parseConfig()
