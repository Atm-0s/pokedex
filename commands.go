package main

import (
	"Atm-0s/pokedex/internal/pokeapi"
	"fmt"
	"math/rand"
	"os"
)

// A struct which holds the info for a command
type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config, *string) error
}

// Command to exit the REPL
func commandExit(config *pokeapi.Config, arg *string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// Command to display the available commands.
func commandHelp(config *pokeapi.Config, arg *string) error {
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
func commandMap(config *pokeapi.Config, arg *string) error {
	if config.Next == nil {
		url := pokeapi.LocationAreasURL
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
func commandMapB(config *pokeapi.Config, arg *string) error {
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

func commandExplore(config *pokeapi.Config, arg *string) error {
	if arg == nil {
		fmt.Println("Missing location area argument")
		fmt.Println("Example usage:")
		fmt.Println("explore pastoria-city-area")
		return nil
	}

	location := *arg
	area, err := config.GetSingleLocationArea(location)
	if err != nil {
		fmt.Println("error fetching location area")
		fmt.Println("invalid location area name")
		return err
	}

	for _, encounters := range area.PokemonEncounters {
		fmt.Println(encounters.Pokemon.Name)
	}
	return nil

}

/*
I need to make a command that attempts to catch a pokemon
I will need to use a random number to determine if the pokemon is caught
I can use the base experience to determine the catch rate
the highest base experience is 608 (apparently)
I could pick a number between 1 and base xp
if the number is below X (100 for example)
pokemon is caught
*/

func commandCatch(config *pokeapi.Config, arg *string) error {
	if arg == nil {
		fmt.Println("Missing pokemon name argument")
		fmt.Println("Example Usage:")
		fmt.Println("catch pikachu")
		return nil
	}
	pokemonName := *arg

	pokemon, err := config.GetPokemonData(pokemonName)
	if err != nil {
		fmt.Println("Error fetching pokemon data")
		fmt.Println("Invalid pokemon name")
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	baseXP := pokemon.BaseExperience
	catchNum := rand.Intn(baseXP + 1)
	if catchNum <= 100 {
		fmt.Printf("%s was caught!\n", pokemonName)
		config.Pokedex[pokemonName] = pokemon
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped\n", pokemonName)
	}

	return nil

}

// I need to make an inspect command that views caught pokemon
func commandInspect(config *pokeapi.Config, arg *string) error {
	if arg == nil {
		fmt.Println("Missing pokemon name argument")
		fmt.Println("Example Usage:")
		fmt.Println("Inspect pikachu")
		return nil
	}

	pokemon, ok := config.Pokedex[*arg]
	if !ok {
		fmt.Printf("%s has not been caught yet!\n", *arg)
	} else {
		p := fmt.Printf
		p("Name: %v\n", pokemon.Name)
		p("Height: %v\n", pokemon.Height)
		p("Weight: %v\n", pokemon.Weight)

		fmt.Println("Stats:")
		for _, s := range pokemon.Stats {
			p("  -%v: %v\n", s.Stat.Name, s.BaseStat)
		}

		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			p("  - %v\n", t.Type.Name)
		}
	}
	return nil

}

func commandPokedex(config *pokeapi.Config, arg *string) error {
	for _, pokemon := range config.Pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
