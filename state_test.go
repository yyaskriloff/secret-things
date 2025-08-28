package main

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

var filePath, _ = filepath.Abs("./state/app.json")

func cleanup() {
	os.Remove(filePath)
}

func TestStateCreateion(t *testing.T) {
	state := appState{}
	state.Load()

	defer cleanup()

	_, err := os.Stat(filePath)
	if errors.Is(err, os.ErrNotExist) {
		t.Error("Failed to create file")
	}

}

func TestEnvMethods(t *testing.T) {
	defer cleanup()
	state := appState{
		Environments: []environment{
			{
				Name: "dev",
				Keys: []string{},
			},
		},
	}
	state.AddEnv("prod")

	if envLength := len(state.Environments); envLength < 2 {
		t.Errorf("The number of envs are %v when it's supposed to be 2", envLength)
	}

	state.RemoveEnv("dev")

	if state.Environments[0].Name != "prod" {
		t.Error("Failed to remove dev env")
	}

}

func TestKeyMethods(t *testing.T) {
	defer cleanup()
	state := appState{
		Environments: []environment{
			{
				Name: "dev",
				Keys: []string{},
			},
		},
	}

	state.AddSecret("dev", "HELLO")

	if keyLength := len(state.Environments[0].Keys); keyLength < 1 {
		t.Errorf("The number of envs keys are %v when it's supposed to be 1", keyLength)
	}

	state.RemoveSecret("dev", "HELLO")
	if keyLength := len(state.Environments[0].Keys); keyLength != 0 {
		t.Errorf("The number of envs keys are %v when it's supposed to be 0", keyLength)
	}

}
