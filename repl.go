package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/fadhilradh/pokedex-cli/config"
)

func StartCLI(config *config.Config) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(color.InGreen(`
Welcome to Pokedex CLI game !

In this game you can : 
> explore locations in the Pokemon world 
> encounter pokemon in each location 
> catch pokemon and inspect them

Type 'help' to see all available commands in this game.

Catch 'em all !

--- made by fadhilradh with Go ---
	`))
	GetInput(scanner, "Pokedex > ", mainCommands, true)
}

func GetInput(scanner *bufio.Scanner, msg string, commandList func() map[string]cliCommand, isWelcome bool) {
	for {
		if isWelcome {
			fmt.Print(color.Red + msg + color.Reset)
		} else {
			fmt.Print(msg)
		}
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
		"map": {
			name:        "map",
			description: "List next 10 locations on the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous 10 locations on the map",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore [location name]",
			description: "Explore a location and find Pokemons",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch [pokemon name]",
			description: "Attempt to catch a Pokemon. Success chance is based on Pokemon's level and rarity",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect [pokemon name]",
			description: "Inspect a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught pokemon",
			callback:    commandPokedex,
		},
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
