package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

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

	for _, cmd := range mainCommands() {
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
	areaName := args[0]
	locDetail, err := cfg.Client.GetLocDetail(areaName)

	if err != nil {
		return err
	}

	// TODO : change to cases.Title
	fmt.Println("\nExploring " + strings.ReplaceAll(strings.Title(areaName), "-", " ") + "... \n")

	rand.NewSource(time.Now().UnixNano())
	randIdx := rand.Intn(len(locDetail.PokemonEncounters))
	pokemon := locDetail.PokemonEncounters[randIdx].Pokemon.Name

	// TODO : change to cases.Title
	fmt.Printf("A wild %s encountered ! \n \n", strings.Title(pokemon))
	fmt.Println(
		`What will you do ?

- battle
- run

[type one of the commands above to continue]
	`)

	scanner := bufio.NewScanner(os.Stdin)
	GetInput(scanner, "vs "+pokemon+" > ", EncounterCommands)

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
	fmt.Printf("You may now inspect it with 'inspect %s' command. \n", pokemon.Name)

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

func commandPokedex(cfg *config.Config, args ...string) error {
	pokemons := cfg.CaughtPokemons
	if len(pokemons) == 0 {
		fmt.Println("You have no pokemon. Catch them !")
		return nil
	}
	fmt.Println("Here is your pokemon collection :")
	for _, v := range cfg.CaughtPokemons {
		fmt.Println("- ", v.Name)
	}

	return nil
}

func commandBattle(cfg *config.Config, params ...string) error {
	fmt.Println("Battle ...")
	return nil
}

func commandRun(cfg *config.Config, params ...string) error {
	fmt.Println("Running away ...")

	return nil
}
