package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	argsLength := len(os.Args)

	if argsLength < 2 {
		log.Fatal("no command was not provided")
	}

	env := flag.String("env", "", "The name of env to assign the vars")
	shouldDelete := flag.Bool("delete", false, "wether the left over vars should be deleted")
	flag.Parse()

	cmd := os.Args[1]

	config := configuration{}
	config.LoadConfig()

	state := appState{}
	state.Load()

	secrets := secretsManager{}
	secrets.Init(&config)

	switch cmd {
	case "env":
		if argsLength < 3 {
			log.Fatal("no env command was provided")
		}
		subCommand := os.Args[2]
		if argsLength < 4 {
			log.Fatal("no env name was provided")
		}

		name := os.Args[3]
		switch subCommand {
		case "create":
			state.AddEnv(name)
		case "delete":
			state.RemoveEnv(name)

		}
	case "set":

		if argsLength < 3 {
			log.Fatal("no var was provided")
		}

		if argsLength < 4 {
			log.Fatal("no var value was provided")
		}

		varName := os.Args[3]
		varValue := os.Args[4]

		if env == nil {
			log.Fatal("you need to set the env flag")
		}

		err := secrets.Set(*env, varName, varValue)

		if err != nil {
			log.Fatal(err)
		}

		state.AddSecret(*env, varName)

	case "load":
		if argsLength < 3 {
			log.Fatal("no file was provided")
		}

		if env == nil {
			log.Fatal("you need to set the env flag")
		}

		fileName := os.Args[3]

		keyValues, err := Parse(fileName)

		if err != nil {
			log.Fatal(err)

		}

		var keysInStore []string

		secrets.ListKeys(*env, &keysInStore, nil)

		for _, k := range keysInStore {

			_, ok := keyValues[k]

			if !ok && *shouldDelete {
				//
				secrets.Remove(*env, k)
				state.RemoveSecret(*env, k)

			}
			// if found make sure it's in state
			if ok {
				state.AddSecret(*env, k)
			}

			// removing so we don't recreate
			delete(keyValues, k)

			// keysInStore = append(keysInStore[:i], keysInStore[i+1:]...)

		}

		// creating
		for k, v := range keyValues {
			secrets.Set(*env, k, v)
			state.AddSecret(*env, k)
		}

	case "sync":
		if argsLength < 3 {
			log.Fatal("no file was provided")
		}

		fileName := os.Args[3]

		keyValues, err := Parse(fileName)

		secrets.GetValues(*env, keyValues, nil)

		if err != nil {
			log.Fatal(err)
		}

		varsForFile := Print(keyValues)

		filePath, _ := filepath.Abs("./state/app.json")
		dirPath, _ := filepath.Abs("./state")

		_, err = os.Stat(dirPath)
		if errors.Is(err, os.ErrNotExist) {
			os.Mkdir(dirPath, os.ModePerm)
		} else if err != nil {
			log.Fatal(err)
		}
		// potential issue if fail to write to file we just trunicated and lost all vars
		f, err := os.Create(filePath)
		if err != nil {
			log.Fatalf("Failed to create/truncate file: %v", err)
		}
		defer f.Close()

		_, err = f.Write(varsForFile.Bytes())
		if err != nil {
			log.Fatalf("Failed to write to file: %v", err)
		}

	default:
		fmt.Printf("%s is not a recognized command", cmd)

	}

	state.Write()

}
