package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	argsLength := len(os.Args)

	if argsLength < 2 {
		log.Fatal("no command was not provided")
	}

	env := flag.String("env", "", "The name of env to assign the vars")
	flag.Parse()

	cmd := os.Args[1]

	config := configuration{}
	config.LoadConfig()

	state := appState{}
	state.Load()

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

		fmt.Printf("Setting %s to %s", varName, varValue)
	case "load":
		if argsLength < 3 {
			log.Fatal("no file was provided")
		}

		fmt.Println(*env)
	case "sync":
		if argsLength < 3 {
			log.Fatal("no file was provided")
		}

		fmt.Println(*env)
	default:
		fmt.Printf("%s is not a recognized command", cmd)

	}

}
