package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {

	wordSlice := strings.Fields(strings.ToLower(text))
	return wordSlice
}

func runREPL() {
	input := bufio.NewScanner(os.Stdin)
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
				err := command.callback()
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
	}
}
