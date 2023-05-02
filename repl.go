package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fadhilradh/pokedex-cli/config"
)

func StartCLI(config *config.Config) {
	scanner := bufio.NewScanner(os.Stdin)

	GetInput(scanner, "Pokedex > ", mainCommands)
}

func GetInput(scanner *bufio.Scanner, title string, commandList func() map[string]cliCommand) {
	fmt.Print(title)
	for {
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		input := scanner.Text()
		words := strings.Fields(input)

		if len(input) == 0 {
			continue
		}

		command, exists := commandList()[words[0]]
		if exists {
			params := []string{}
			if len(words) > 1 {
				params = words[1:]
			}
			err := command.callback(&cfg, params...)
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

func mainCommands() map[string]cliCommand {
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
		"catch": {
			name:        "catch",
			description: "Catch 'em all !",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "list all caught pokemon",
			callback:    commandPokedex,
		},
	}
}

func EncounterCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"battle": {
			name:        "battle",
			description: "Battle Pokemon with chance to capture it",
			callback:    commandBattle,
		},
		"run": {
			name:        "run",
			description: "Run away with your life !",
			callback:    commandRun,
		},
	}
}
