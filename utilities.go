package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Represents our config file
type Config struct {
	Secrets struct {
		Watch2gether string `json:"w2g"`
		Discord      string `json:"discord"`
	} `json:"secrets"`
}

/**
 * @brief Handles reading our config file and returning
 * a Config object that represents the file's contents.
 *
 * @param file Represents the filepath to config.json
 */
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
