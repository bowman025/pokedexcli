package main

import (
	"fmt"
	"os"
)

func commandExit(conf *config, param string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp(conf *config, param string) error {
	_, err := fmt.Printf("Welcome to the Pokedex!\nUsage:\n")
	if err != nil {
		return err
	}

	for _, command := range conf.commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandMap(conf *config, param string) error {
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

func commandMapB(conf *config, param string) error {
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

func commandExplore(conf *config, param string) error {
	return nil
}
