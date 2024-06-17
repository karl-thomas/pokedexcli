package main

import (
	"time"

	"github.com/karl-thomas/pokedexcli/pokeapi"
	"github.com/karl-thomas/pokedexcli/pokecache"
)

type config struct {
	pokeApi             pokeapi.Client
	nextLocationAreaUrl *string
	prevLocationAreaUrl *string
	cache               *pokecache.Cache
}

func main() {
	config := config{
		pokeApi: pokeapi.NewClient(),
		cache:   pokecache.NewCache(int64(time.Duration(30).Seconds())),
	}
	StartRepl(&config)
}
