package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2/"

func FetchLocationArea() {
	client := NewClient()
	resp, err := client.httpClient.Get(baseURL + "location-area")
	if err != nil {
		log.Fatal(err)
	}
	var locationAreaResponse LocationAreaResponse
	handler(resp, &locationAreaResponse)
	log.Println(locationAreaResponse)
}

func handler(r *http.Response, params interface{}) {
	defer r.Body.Close()

	dat, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(dat, &params)
	if err != nil {
		return
	}
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

type LocationAreaResponse struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}
type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
