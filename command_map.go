package main

import (
	"fmt"

	"github.com/karl-thomas/pokedexcli/pokeapi"
)

func commandMap() error {
	client := pokeapi.NewClient()
	resp, err := client.FetchLocationAreas()
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil
}
