package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/fadhilradh/pokedex-cli/internal/pokeapi"
)

type config struct {
	Client     pokeapi.Client
	NextLocURL *string
	PrevLocURL *string
}

var cfg = config{}

func Start(config *config) {

	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		commands := GetCommand()
		commands[scanner.Text()].callback()
	}
}
