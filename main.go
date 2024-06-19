package main

import (
	"github.com/karl-thomas/pokedexcli/pokeapi"
)

type config struct {
	pokeApi             pokeapi.Client
	nextLocationAreaUrl *string
	prevLocationAreaUrl *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {
	config := config{
		pokeApi:       pokeapi.NewClient(),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	StartRepl(&config)
}
