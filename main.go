package main

func main() {
	conf := config{
		commands:    getCommands(),
		callPokeApi: getPokeResponse,
	}
	startRepl(&conf)
}
