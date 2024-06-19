package main

import (
	"errors"
	"fmt"
)

func commandInspect(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: catch <pokemon>")
	}

	pokemonName := args[0]
	pokemon, ok := config.caughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("You haven't caught %s yet", pokemonName)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf("%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Abilities:")
	for _, ability := range pokemon.Abilities {
		fmt.Printf(" - %s\n", ability.Ability.Name)
	}
	fmt.Println("Types:")
	for _, typeEntry := range pokemon.Types {
		fmt.Printf(" - %s\n", typeEntry.Type.Name)
	}

	return nil
}
