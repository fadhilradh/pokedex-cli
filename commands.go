package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func commandHelp(cfg *config, params ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommand() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()

	return nil
}

func commandExit(cfg *config, params ...string) error {
	os.Exit(0)

	return nil
}

func commandMap(cfg *config, params ...string) error {
	nextUrl := cfg.NextLocURL
	maps, err := cfg.Client.ListLocations(nextUrl)

	if err != nil {
		return err
	}

	for _, loc := range maps.Results {
		fmt.Println(loc.Name)
	}
	fmt.Println()

	cfg.NextLocURL = maps.Next
	cfg.PrevLocURL = maps.Previous

	return nil
}

func commandMapBack(cfg *config, params ...string) error {
	prevUrl := cfg.PrevLocURL
	if prevUrl == nil {
		fmt.Println("Oops. There is no previous map")
	} else {
		maps, err := cfg.Client.ListLocations(prevUrl)
		if err != nil {
			return err
		}
		for _, loc := range maps.Results {
			fmt.Println(loc.Name)
		}
		fmt.Println()

		cfg.NextLocURL = maps.Next
		cfg.PrevLocURL = maps.Previous
	}

	return nil
}

func commandExplore(cfg *config, args ...string) error {
	fmt.Println(args[0])
	locDetail, err := cfg.Client.GetLocDetail(args[0])

	if err != nil {
		return err
	}

	fmt.Println("Exploring " + args[0])

	for _, data := range locDetail.PokemonEncounters {
		fmt.Println("- " + data.Pokemon.Name)
	}

	return nil
}
