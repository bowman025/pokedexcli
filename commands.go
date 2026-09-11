package main

import (
	"fmt"
	"os"
)

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
	pokeRes, err := conf.pokeApiClient.GetPokeResponse(conf.nextLocationsURL)
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

	pokeRes, err := conf.pokeApiClient.GetPokeResponse(conf.prevLocationsURL)
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
