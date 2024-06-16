package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2/"

func (c *Client) FetchLocationAreas() (LocationAreaResponse, error) {
	resp, err := c.httpClient.Get(baseURL + "location-area")
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
