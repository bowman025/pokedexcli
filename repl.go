package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type config struct {
	commands         map[string]cliCommand
	nextLocationsURL *string
	prevLocationsURL *string
	callPokeApi      func(*string) (pokeResponse, error)
}

type pokeResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 location areas",
			callback:    commandMapB,
		},
	}
}

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp(conf *config) error {
	_, err := fmt.Printf("Welcome to the Pokedex!\nUsage:\n")
	if err != nil {
		return err
	}

	for _, command := range conf.commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandMap(conf *config) error {
	pokeRes, err := conf.callPokeApi(conf.nextLocationsURL)
	if err != nil {
		return err
	}

	conf.nextLocationsURL = pokeRes.Next
	conf.prevLocationsURL = pokeRes.Previous

	for _, locArea := range pokeRes.Results {
		fmt.Println(locArea.Name)
	}

	return nil
}

func commandMapB(conf *config) error {
	if conf.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	pokeRes, err := conf.callPokeApi(conf.prevLocationsURL)
	if err != nil {
		return err
	}

	conf.nextLocationsURL = pokeRes.Next
	conf.prevLocationsURL = pokeRes.Previous

	for _, locArea := range pokeRes.Results {
		fmt.Println(locArea.Name)
	}

	return nil
}

func startRepl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Printf("Error reading input: %v\n", err)
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
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback(conf)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func cleanInput(text string) []string {
	lowText := strings.ToLower(text)
	return strings.Fields(lowText)
}

func getPokeResponse(urlAddress *string) (pokeResponse, error) {
	urlValue := "https://pokeapi.co/api/v2/location-area"
	if urlAddress != nil {
		urlValue = *urlAddress
	}

	res, err := http.Get(urlValue)
	if err != nil {
		return pokeResponse{}, fmt.Errorf("API error: %v", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return pokeResponse{}, fmt.Errorf("error reading the request: %v", err)
	}

	pokeRes := pokeResponse{}
	err = json.Unmarshal(data, &pokeRes)
	if err != nil {
		return pokeResponse{}, fmt.Errorf("error during unmarshal: %v", err)
	}

	return pokeRes, nil
}
