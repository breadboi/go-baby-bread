package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config Represents our config file
type Config struct {
	Secrets struct {
		Watch2gether string `json:"w2g"`
		Discord      string `json:"discord"`
	} `json:"secrets"`
}

// LoadConfiguration Handles reading our config file and returning a Config object that represents the file's contents.
func LoadConfiguration(file string) Config {
	var config Config

	configFile, err := os.Open(file)

	defer configFile.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	jsonParser := json.NewDecoder(configFile)

	jsonParser.Decode(&config)

	return config
}
