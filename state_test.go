package main

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestStateCreateion(t *testing.T) {
	state := appState{}
	state.Load()

	filePath, _ := filepath.Abs("./state/app.json")

	_, err := os.Stat(filePath)
	if errors.Is(err, os.ErrNotExist) {
		t.Fail()
	}

	os.Remove(filePath)

}

