package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bowman025/pokedexcli/internal/pokeapi"
)

type config struct {
	commands         map[string]cliCommand
	nextLocationsURL *string
	prevLocationsURL *string
	pokeApiClient    pokeapi.Client
	pokedex          map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:        "catch",
			description: "Try to catch a Pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Explore the selected location",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous 20 location areas",
			callback:    commandMapB,
		},
	}
}

func startRepl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Printf("error reading input: %v\n", err)
			}
			break
		}

		input := scanner.Text()
		clean := cleanInput(input)
		if len(clean) == 0 {
			continue
		}

		commandName := clean[0]
		cmd, exists := conf.commands[commandName]
		if !exists {
			fmt.Println("unknown command")
			continue
		}

		args := []string{}
		if len(clean) > 1 {
			args = clean[1:]
		}
		err := cmd.callback(conf, args...)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
	}
}

func cleanInput(text string) []string {
	lowText := strings.ToLower(text)
	return strings.Fields(lowText)
}
