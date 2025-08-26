package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	argsLength := len(os.Args)

	if argsLength < 2 {
		panic("no command was not provided")
	}

	env := flag.String("env", "", "The name of env to assign the vars")
	flag.Parse()

	cmd := os.Args[1]

	switch cmd {
	case "config":
		if argsLength < 3 {
			panic("no config command was provided")

		}
		subCommand := os.Args[2]
		if argsLength < 4 {
			panic("no app name was provided")
		}
		// name := os.Args[3]

		switch subCommand {
		case "app":
			fmt.Println("app subcommand")
		case "profile":
			fmt.Println("profile subcommand")
		default:
			panic(fmt.Sprintf("%s is not a recognized command", subCommand))
		}

	case "env":
		if argsLength < 3 {
			panic("no env command was provided")
		}
		subCommand := os.Args[2]
		if argsLength < 4 {
			panic("no env name was provided")
		}

		// name := os.Args[3]
		switch subCommand {
		case "create":
			fmt.Println("create subcommand")
		case "delete":
			fmt.Println("delete subcommand")

		}
	case "set":

		if argsLength < 3 {
			panic("no var was provided")
		}

		if argsLength < 4 {
			panic("no var value was provided")
		}

		varName := os.Args[3]
		varValue := os.Args[4]

		fmt.Printf("Setting %s to %s", varName, varValue)

	case "load":
		if argsLength < 3 {
			panic("no file was provided")
		}

		fmt.Println(*env)

	case "sync":
		if argsLength < 3 {
			panic("no file was provided")
		}

		fmt.Println(*env)

	default:
		fmt.Printf("%s is not a recognized command", cmd)

	}

}
