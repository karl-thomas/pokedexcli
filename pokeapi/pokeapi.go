package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/karl-thomas/pokedexcli/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2/"

func (c *Client) FetchLocationAreas(pageUrl *string) (LocationAreaResponse, error) {
	fullUrl := baseURL + "location-area"
	if pageUrl != nil {
		fullUrl = *pageUrl
	}
	if cached, ok := c.cache.Get(fullUrl); ok {
		var locationAreaResponse LocationAreaResponse
		err := json.Unmarshal(cached, &locationAreaResponse)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		fmt.Println("Cache hit")
		return locationAreaResponse, nil
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

	c.cache.Set(fullUrl, dat)
	return locationAreaResponse, nil
}

func (c *Client) GetLocation(locationAreaName string) (LocationArea, error) {
	fullUrl := baseURL + "location-area/" + locationAreaName
	if cached, ok := c.cache.Get(fullUrl); ok {
		var locationArea LocationArea
		err := json.Unmarshal(cached, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		fmt.Println("Cache hit")
		return locationArea, nil
	}

	resp, err := c.httpClient.Get(fullUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}
	var locationArea LocationArea
	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Set(fullUrl, dat)
	return locationArea, nil
}

type Client struct {
	cache      *pokecache.Cache
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: pokecache.NewCache(time.Second * 10),
	}
}
