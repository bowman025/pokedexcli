package main

import (
	"time"

	"github.com/bowman025/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	conf := config{
		commands:      getCommands(),
		pokeApiClient: pokeClient,
	}
	startRepl(&conf)
}
