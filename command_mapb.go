package main

import (
	"errors"
	"fmt"
)

func commandMapb(config *config) error {
	if config.prevLocationAreaUrl == nil {
		return errors.New("you are on the first page")
	}
	resp, err := config.pokeApi.FetchLocationAreas(config.prevLocationAreaUrl)
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
