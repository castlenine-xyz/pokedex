package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner((os.Stdin))

	for {
		//print prompt
		fmt.Print("prompt>")

		//scan prompt
		scanner.Scan()
		text := scanner.Text()

		//clean text from cli
		cleaned := cleanInput(text)
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}
		// hit enter condition make a new line
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
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}

}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "brings you to help hell",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "lists some location areas, type again to page forward",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "page back through location list",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "lists pokemon in a location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon}",
			description: "trys to catch a specifed pokemon and add it to your pokedex",
			callback:    callbackCatch,
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
