package main

import (
	"github.com/fadhilradh/pokedex-cli/internal/pokeapi"
)

type config struct {
	Client         pokeapi.Client
	NextLocURL     *string
	PrevLocURL     *string
	CaughtPokemons map[string]pokeapi.Pokemon
}
