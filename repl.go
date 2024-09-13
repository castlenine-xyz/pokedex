package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner((os.Stdin))

	for {
		//print prompt
		fmt.Print("prompt>")

		//scan prompt
		scanner.Scan()
		text := scanner.Text()

		//clean text from cli
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		//parsing
		commandName := cleaned[0]
		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		// invalid command check
		if !ok {
			fmt.Printf("Invald command: %s\n", commandName)
			continue
		}
		command.callback()
	}

}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "brings you to help hell",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "kills program, same as ctrl c",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	//tokenizer and washer
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
