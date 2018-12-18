package config

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"os"
)

var (
	Logger *zap.Logger
)

type Config struct {
	Port    string
	Name    string
	Version string
}

func Parser() Config {
	file, _ := os.Open("./conf.json")
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(configuration.Port)

	return configuration
}
