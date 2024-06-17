package main

import (
	"errors"
	"fmt"
)

func commandExplore(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("explore command requires a location area name")
	}
	locationAreaName := args[0]
	resp, err := config.pokeApi.GetLocation(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in %s:\n", locationAreaName)
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
