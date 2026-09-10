package main

import (
	"github.com/bowman025/pokedexcli/internal/pokeapi"
)

func main() {
	conf := config{
		commands:    getCommands(),
		callPokeApi: pokeapi.GetPokeResponse,
	}
	startRepl(&conf)
}
