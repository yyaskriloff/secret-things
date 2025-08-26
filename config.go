package main

import (
	"encoding/json"
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
		panic("No config file found")
	}

	err = json.Unmarshal([]byte(contents), &c)

	if err != nil {
		panic("Unable to parse config file")
	}

}
