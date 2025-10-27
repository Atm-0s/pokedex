package main

import (
	"Atm-0s/pokedex/internal/pokeapi"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Take Smiley's advice and go through and comment everything!

// Cleans the input to only use the first word in lower case
func cleanInput(text string) []string {

	wordSlice := strings.Fields(strings.ToLower(text))
	return wordSlice
}

/*
The main logic for the Read, Evaluate, Print, Loop
Creates config which creates a client which creates a cache
inf loop
reads the input and handle errors
grabs the command and executes
loop back
*/

func runREPL() {
	input := bufio.NewScanner(os.Stdin)
	config := &pokeapi.Config{
		PClient: pokeapi.NewClient(5*time.Minute, 10*time.Second),
		Pokedex: make(map[string]pokeapi.Pokemon),
	}

	for {
		fmt.Print("Pokedex > ")
		if input.Scan() {
			text := input.Text()
			cleanText := cleanInput(text)
			if len(cleanText) == 0 {
				continue
			}
			commandName := cleanText[0]
			var commandArgument *string
			if len(cleanText) > 1 {
				commandArgument = &cleanText[1]
			}

			command, ok := getCommands()[commandName]
			if !ok {
				fmt.Println("Unknown command")
				continue
			} else {
				err := command.callback(config, commandArgument)
				if err != nil {
					fmt.Print(err)
				}
				continue
			}

		}
	}
}

// A map of the commands with descriptions
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex.",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays a list of 20 location areas per page. Multiple uses will display the next 20.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations listed paged through by map command.",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Takes a location area as an argument and displays a list of pokemon encountered in that area.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon and add it to the pokedex.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "inspect",
			description: "List the pokemon in your pokedex",
			callback:    commandPokedex,
		},
	}
}
