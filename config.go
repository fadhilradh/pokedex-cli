package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Config struct {
	NextLocURL *string
	PrevLocURL *string
}

var Cfg = Config{}

func Start(config *Config) {

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
