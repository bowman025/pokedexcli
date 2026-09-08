package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		clean := cleanInput(input)
		fmt.Printf("Your command was: %s\n", clean[0])
		scanner.Err()
	}
}
