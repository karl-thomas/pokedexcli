package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)
	replMessage()
	for scanner.Scan() {
		inputs := cleanInput(scanner.Text())
		if len(inputs) == 0 {
			replMessage()
			continue
		}

		command := inputs[0]
		args := []string{}
		if len(inputs) > 1 {
			args = inputs[1:]
		}

		if res, exists := commands()[command]; exists {
			err := res.callback(config, args...)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command:", command)
		}
		fmt.Println()
		replMessage()
	}
}

func cleanInput(input string) []string {
	return strings.Fields(strings.TrimSpace(strings.ToLower(input)))
}

func replMessage() {
	fmt.Print("Pokedex > ")
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "The map command displays the names of 20 location areas in the Pokemon world. Each subsequent call to map should display the next 20 locations, and so on.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "The map command displays the names of previous 20 location areas in the Pokemon world. This is the reverse of the map command.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location area name>",
			description: "The explore command displays the names of the Pokemon in a location area. The location area is specified by the user.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "The catch command attempts to catch a Pokemon. The Pokemon is specified by the user.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "The inspect command displays the details of a Pokemon. The Pokemon is specified by the user.",
			callback:    commandInspect,
		},
	}
}
