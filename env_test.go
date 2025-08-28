package main

import (
	"io"
	"testing"
)

func TestEnvParsing(t *testing.T) {
	filepath := "./.env"
	vars, err := Parse(filepath)
	if err != nil && err != io.EOF {
		t.Fatalf("Failed with the error %v", err)
	}

	if len(vars) < 2 {
		t.Fatal("Result doesn't have two values")
	}

	t.Logf("value 1: %s", vars["VALUE1"])
	t.Logf("value 2: %s", vars["VALUE2"])

}

func TestEnvPrint(t *testing.T) {
	vars := map[string]string{
		"VALUE1": "HELLO",
		"VALUE2": "WORLD",
	}

	WriteEnv(vars, "./test.env")

}
