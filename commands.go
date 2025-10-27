package main

import (
	"Atm-0s/pokedex/internal/pokeapi"
	"fmt"
	"os"
)

// A struct which holds the info for a command
type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config) error
}

// Command to exit the REPL
func commandExit(config *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// Command to display the available commands.
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

// Command that uses the pokeapi.Config.GetLocationAreas method to
// displat a list of 20 locations area names
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

// As above but desplays the previous 20 pages a la pagination style
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
