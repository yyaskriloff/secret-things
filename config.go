package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type configuration struct {
	App     string `json:"app"`
	Region  string `json:"region"`
	Profile string `json:"profile"`
}

func (c *configuration) LoadConfig() {
	filePath, _ := filepath.Abs("./secrets.config.json")
	contents, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("No config file found")
	}

	err = json.Unmarshal([]byte(contents), &c)

	if err != nil {
		log.Fatal("Unable to parse config file")
	}

}
