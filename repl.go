package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	replMessage()
	for scanner.Scan() {
		command := cleanInput(scanner.Text())

		if res, exists := commands()[command]; exists {
			err := res.callback()
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

func cleanInput(input string) string {
	return strings.TrimSpace(strings.ToLower(input))
}

func replMessage() {
	fmt.Print("Pokedex > ")
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	next string
	prev string
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
			callback:    commandMap,
		},
	}
}
