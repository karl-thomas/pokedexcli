package main

import (
	"fmt"
)

func commandMap(config *config) error {
	resp, err := config.pokeApi.FetchLocationAreas(config.nextLocationAreaUrl)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")
	for _, locationArea := range resp.Results {
		fmt.Printf(" - %s\n", locationArea.Name)
	}
	config.nextLocationAreaUrl = resp.Next
	config.prevLocationAreaUrl = resp.Previous
	return nil
}
