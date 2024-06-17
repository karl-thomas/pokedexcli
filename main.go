package main

import (
	"github.com/karl-thomas/pokedexcli/pokeapi"
)

type config struct {
	pokeApi             pokeapi.Client
	nextLocationAreaUrl *string
	prevLocationAreaUrl *string
}

func main() {
	config := config{
		pokeApi: pokeapi.NewClient(),
	}
	StartRepl(&config)
}
