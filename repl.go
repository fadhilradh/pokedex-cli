package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func StartCLI(config *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		input := scanner.Text()
		words := strings.Fields(input)

		command, exists := getCommand()[words[0]]
		if exists {
			locName := []string{}
			if len(words) > 1 {
				locName = words[1:]
			}
			err := command.callback(&cfg, locName...)
			if err != nil {
				log.Fatal(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}

}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List all locations of PokeMap",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List all locations of PokeMap",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "List all pokemons in a location",
			callback:    commandExplore,
		},
	}

}
