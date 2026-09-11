package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(conf *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp(conf *config, args ...string) error {
	_, err := fmt.Printf("Welcome to the Pokedex!\nUsage:\n")
	if err != nil {
		return err
	}

	for _, command := range conf.commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandMap(conf *config, args ...string) error {
	pokeRes, err := conf.pokeApiClient.GetLocationList(conf.nextLocationsURL)
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

func commandMapB(conf *config, args ...string) error {
	if conf.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	pokeRes, err := conf.pokeApiClient.GetLocationList(conf.prevLocationsURL)
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

func commandExplore(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("location name not provided")
	}

	locationName := args[0]
	location, err := conf.pokeApiClient.GetLocation(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", location.Name)

	if len(location.PokemonEncounters) == 0 {
		fmt.Println("Did not find any Pokemon")
	} else {
		fmt.Println("Found Pokemon:")
		for _, pokemon := range location.PokemonEncounters {
			fmt.Printf("- %s\n", pokemon.Pokemon.Name)
		}
	}

	return nil
}
