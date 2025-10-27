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

func cleanInput(text string) []string {

	wordSlice := strings.Fields(strings.ToLower(text))
	return wordSlice
}

func runREPL() {
	input := bufio.NewScanner(os.Stdin)
	config := &pokeapi.Config{
		PClient: pokeapi.NewClient(5*time.Minute, 5*time.Second),
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

			command, ok := getCommands()[commandName]
			if !ok {
				fmt.Println("Unknown command")
				continue
			} else {
				err := command.callback(config)
				if err != nil {
					fmt.Print(err)
				}
				continue
			}

		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
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
	}
}
