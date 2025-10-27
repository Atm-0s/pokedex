package pokeapi

import (
	"Atm-0s/pokedex/internal/pokecache"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Create a new pokeapi.Client struct using inputted timeout and interval durations
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		client: http.Client{
			Timeout: timeout,
		},
	}
}

// *Config method that takes a url as input and returns a struct of location areas
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

	// Add the data to the cache then return it
	c.PClient.cache.Add(url, data)
	return locations, nil
}

/*
I need to make an 'explore' command that accepts a location name
and prints a list of pokemon found in the area.
The GetLocationAreas function returns the LocationAreas struct.
The Results field of LocationAreas is a slice of anonymous structs
containing a Name and URL field which themselves are jsons of a
single LocationArea name and LocationArea url respectively.

The LocationArea type on PokeAPI contains a field for pokemon_encounters.
*/

// This function retrives a single locationArea and unmarshals into a LocationArea struct
func (c *Config) GetSingleLocationArea(locationAreaName string) (LocationArea, error) {
	url := LocationAreasURL + fmt.Sprintf("/%s", locationAreaName)

	// Check the Cache for the data
	if val, ok := c.PClient.cache.Get(url); ok {
		location := LocationArea{}
		err := json.Unmarshal(val, &location)
		if err != nil {
			return LocationArea{}, err
		}
		return location, nil
	}

	// If no cache data found, make the request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	// Handle the response
	resp, err := c.PClient.client.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	// Handle the data
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	location := LocationArea{}
	err = json.Unmarshal(data, &location)
	if err != nil {
		return LocationArea{}, err
	}

	// Add the data to the cache then return it
	c.PClient.cache.Add(url, data)
	return location, nil

}

func (c *Config) GetPokemonData(pokemonName string) (Pokemon, error) {
	url := pokemonURL + fmt.Sprintf("/%s", pokemonName)

	// Check the Cache for the data
	if val, ok := c.PClient.cache.Get(url); ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	// If no cache data found, make the request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	// Handle the response
	resp, err := c.PClient.client.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	// Handle the data
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	// Add the data to the cache then return it
	c.PClient.cache.Add(url, data)
	return pokemon, nil
}
