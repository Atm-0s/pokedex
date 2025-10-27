package pokeapi

import (
	"Atm-0s/pokedex/internal/pokecache"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	BaseURL = "https://pokeapi.co/api/v2"
)

type Config struct {
	PClient  Client
	Next     *string
	Previous *string
}

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Client struct {
	client http.Client
	cache  *pokecache.Cache
}

// Generic Get request wrapper
func Get(urlPath *string) ([]byte, error) {
	response, err := http.Get(*urlPath)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetLocationAreas(url *string) (LocationAreas, error) {
	jsonData, err := Get(url)
	if err != nil {
		return LocationAreas{}, err
	}

	var locations LocationAreas
	if err := json.Unmarshal(jsonData, &locations); err != nil {
		return LocationAreas{}, err
	}

	return locations, nil

}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		client: http.Client{
			Timeout: timeout,
		},
	}
}
