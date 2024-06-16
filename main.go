package main

import (
	"fmt"
	"os"

	"github.com/karl-thomas/pokedexcli/pokeapi"
)

func main() {
	StartRepl()
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range commands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

type config struct {
	next string
	prev string
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func commandMap() error {
	pokeapi.FetchLocationArea()
	return nil
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
