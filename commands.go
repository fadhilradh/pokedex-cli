package main

import (
	"fmt"
	"os"

	"github.com/fadhilradh/pokedex-cli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommand() {
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
	nextUrl := cfg.NextLocURL
	maps := pokeapi.GetMap(nextUrl)
	for _, loc := range maps.Results {
		fmt.Println(loc.Name)
	}
	fmt.Println()

	cfg.NextLocURL = maps.Next
	cfg.PrevLocURL = maps.Previous

	return nil
}

func commandMapBack() error {
	prevUrl := cfg.PrevLocURL
	if prevUrl == nil {
		fmt.Println("Oops. There is no previous map")
	} else {
		maps := pokeapi.GetMap(prevUrl)
		for _, loc := range maps.Results {
			fmt.Println(loc.Name)
		}
		fmt.Println()

		cfg.NextLocURL = maps.Next
		cfg.PrevLocURL = maps.Previous
	}

	return nil
}
