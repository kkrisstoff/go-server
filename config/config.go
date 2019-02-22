package config

import (
	"encoding/json"
	"os"
)

const configFile = "config/config.json"

type Config struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

func GetConfig() (Config, error) {
	var config Config
	configFile, err := os.Open(configFile)
	defer configFile.Close()
	if err != nil {
		return config, err
	}

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	return config, err
}
