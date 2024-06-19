package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	fullUrl := baseURL + "pokemon/" + pokemonName
	if cached, ok := c.cache.Get(fullUrl); ok {
		var pokemon Pokemon
		err := json.Unmarshal(cached, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		fmt.Println("Cache hit")
		return pokemon, nil
	}

	resp, err := c.httpClient.Get(fullUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	var pokemon Pokemon
	err = json.Unmarshal(dat, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Set(fullUrl, dat)
	return pokemon, nil
}
