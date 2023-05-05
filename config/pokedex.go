package config

import (
	"github.com/fadhilradh/pokedex-cli/internal/pokeapi"
)

type Config struct {
	Client                    pokeapi.Client
	NextLocURL                *string
	PrevLocURL                *string
	CaughtPokemons            map[string]pokeapi.Pokemon
	CurrentPokemonEncountered pokeapi.Pokemon
}
