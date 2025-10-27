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

func (c *Config) GetLocationAreas(pageURL *string) (LocationAreas, error) {
	var url string
	if pageURL != nil {
		url = *pageURL
	}

	// Check the Cache for the data
	if val, ok := c.PClient.cache.Get(url); ok {
		locations := LocationAreas{}
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return LocationAreas{}, err
		}
		return locations, nil
	}

	// If no cache data found, make the request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	// Handle the response
	resp, err := c.PClient.client.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}
	defer resp.Body.Close()

	// Handle the data
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	locations := LocationAreas{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationAreas{}, err
	}

	c.PClient.cache.Add(url, data)
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
