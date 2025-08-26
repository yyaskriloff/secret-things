package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	argsLength := len(os.Args)

	if argsLength < 2 {
		fmt.Println("no command was not provided")
		os.Exit(1)
	}

	env := flag.String("env", "", "The name of env to assign the vars")
	flag.Parse()

	cmd := os.Args[1]

	switch cmd {
	case "config":
		if argsLength < 3 {
			fmt.Println("no config command was provided")
			os.Exit(1)
		}
		subCommand := os.Args[2]
		if argsLength < 4 {
			fmt.Println("no app name was provided")
			os.Exit(1)
		}
		// name := os.Args[3]

		switch subCommand {
		case "app":
			fmt.Println("app subcommand")
		case "profile":
			fmt.Println("profile subcommand")
		default:
			fmt.Printf("%s is not a recognized command", subCommand)
			os.Exit(1)
		}

	case "env":
		if argsLength < 3 {
			fmt.Println("no env command was provided")
			os.Exit(1)
		}
		subCommand := os.Args[2]
		if argsLength < 4 {
			fmt.Println("no env name was provided")
			os.Exit(1)
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
			fmt.Println("no var was provided")
			os.Exit(1)
		}

		if argsLength < 4 {
			fmt.Println("no var value was provided")
			os.Exit(1)
		}

		varName := os.Args[3]
		varValue := os.Args[4]

		fmt.Printf("Setting %s to %s", varName, varValue)

	case "load":
		if argsLength < 3 {
			fmt.Println("no file was provided")
			os.Exit(1)
		}

		fmt.Println(*env)

	case "sync":
		if argsLength < 3 {
			fmt.Println("no file was provided")
			os.Exit(1)
		}

		fmt.Println(*env)

	default:
		fmt.Printf("%s is not a recognized command", cmd)

	}

}
