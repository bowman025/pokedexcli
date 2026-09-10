package main

func main() {
	conf := config{
		commands: getCommands(),
	}
	startRepl(&conf)
}
