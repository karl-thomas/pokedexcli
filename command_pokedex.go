package main

import (
	"fmt"
)

func commandPokedex(config *config, args ...string) error {
	if len(config.caughtPokemon) == 0 {
		fmt.Println("You haven't caught any Pokemon yet!")
		return nil
	}

	fmt.Println("Pokedex:")
	for pokemon := range config.caughtPokemon {
		fmt.Printf("- %s", pokemon)
	}

	return nil
}
