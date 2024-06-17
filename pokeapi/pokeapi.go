package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2/"

func (c *Client) FetchLocationAreas(pageUrl *string) (LocationAreaResponse, error) {
	fullUrl := baseURL + "location-area"
	if pageUrl != nil {
		fullUrl = *pageUrl
	}
	resp, err := c.httpClient.Get(fullUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	var locationAreaResponse LocationAreaResponse
	err = json.Unmarshal(dat, &locationAreaResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	return locationAreaResponse, nil
}

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
