package cli

import (
	"fmt"
	"os"

	"github.com/fadhilradh/pokedex-cli/pokedex"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func GetCommand() map[string]cliCommand {
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
			name:        "map back",
			description: "List all locations of PokeMap",
			callback:    commandMapBack,
		},
	}

}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range GetCommand() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()

	return nil
}

func commandExit() error {
	os.Exit(0)

	return nil
}

func commandMap() error {
	cfg := pokedex.Cfg
	nextUrl := cfg.NextLocURL
	maps := pokedex.GetMap(nextUrl)
	for _, loc := range maps.Results {
		fmt.Println(loc.Name)
	}
	fmt.Println()

	cfg.NextLocURL = maps.Next
	cfg.PrevLocURL = maps.Previous

	return nil
}

func commandMapBack() error {
	maps := pokedex.GetMap(pokedex.Cfg.PrevLocURL)
	for _, loc := range maps.Results {
		fmt.Println(loc.Name)
	}
	fmt.Println()

	pokedex.Cfg.NextLocURL = maps.Next
	pokedex.Cfg.PrevLocURL = maps.Previous

	return nil
}
