package main

import (
	"Atm-0s/pokedex/internal/pokeapi"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config) error
}

func commandExit(config *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *pokeapi.Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandMap(config *pokeapi.Config) error {
	if config.Next == nil {
		url := pokeapi.BaseURL + "/location-area"
		config.Next = &url
	}
	locations, err := config.GetLocationAreas(config.Next)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	config.Next = locations.Next
	config.Previous = locations.Previous
	return nil
}

func commandMapB(config *pokeapi.Config) error {
	if config.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	locations, err := config.GetLocationAreas(config.Previous)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	config.Next = locations.Next
	config.Previous = locations.Previous
	return nil
}
