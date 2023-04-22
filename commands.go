package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"

	"github.com/fadhilradh/pokedex-cli/config"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config.Config, ...string) error
}

func commandHelp(cfg *config.Config, params ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommand() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()

	return nil
}

func commandExit(cfg *config.Config, params ...string) error {
	os.Exit(0)

	return nil
}

func commandMap(cfg *config.Config, params ...string) error {
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

func commandMapBack(cfg *config.Config, params ...string) error {
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

func commandExplore(cfg *config.Config, args ...string) error {
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

func commandCatch(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	pokemon, err := cfg.Client.GetPokemon(args[0])

	if err != nil {
		return err
	}

	chance := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if chance > 40 {
		fmt.Printf("%s escaped!", pokemon.Name)
		fmt.Println()
		return nil
	}

	cfg.CaughtPokemons[pokemon.Name] = pokemon
	fmt.Printf("%s was caught ! \n", pokemon.Name)
	fmt.Println("Congrats ! Here is your pokemon collection :")
	for _, v := range cfg.CaughtPokemons {
		fmt.Println("- ", v.Name)
	}

	return nil
}

func commandInspect(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemon, exist := cfg.CaughtPokemons[args[0]]
	if !exist {
		fmt.Printf("You have not caught %s yet", args[0])
		fmt.Println()
		return nil
	}
	fmt.Printf("Name : %s \n", pokemon.Name)
	fmt.Printf("Weight : %d kg \n ", pokemon.Weight)
	fmt.Printf("Height : %d cm \n", pokemon.Height)

	fmt.Println("Stats")
	for _, stat := range pokemon.Stats {
		fmt.Printf("%s : %d \n", stat.Stat.Name, stat.BaseStat)
	}

	return nil

}
