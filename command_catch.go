package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: catch <pokemon>")
	}

	pokemonName := args[0]

	pokemon, err := config.pokeApi.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	difficulty := rand.Intn(pokemon.BaseExperience)
	if difficulty > threshold {
		return fmt.Errorf("You failed to catch %s", pokemonName)
	}

	config.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("You caught %s!", pokemonName)
	return nil
}
