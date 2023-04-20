package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

		command, exists := getCommand()[input]
		if exists {
			err := command.callback()
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
	}

}
