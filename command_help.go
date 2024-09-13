package main

import "fmt"

func callbackHelp() error {
	fmt.Println("")
	fmt.Println("Welcome to help hell, we don't know either")
	fmt.Printf("heres the manual: \n\n")

	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf(" - %s: %s\n ", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}
