package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type config struct {
	App     string `json:"app"`
	Region  string `json:"region"`
	Profile string `json:"profile"`
}

func GetConfig() config {
	filePath, _ := filepath.Abs("./secrets.config.json")
	contents, err := os.ReadFile(filePath)
	if err != nil {
		panic("No config file found")
	}

	parsedConfig := config{}

	err = json.Unmarshal([]byte(contents), &parsedConfig)

	if err != nil {
		panic("Unable to parse config file")
	}

	return parsedConfig

}
