package main

import (
	"time"

	"github.com/fadhilradh/pokedex-cli/config"
	"github.com/fadhilradh/pokedex-cli/internal/pokeapi"
)

var pokeClient = pokeapi.NewClient(10*time.Second, time.Hour*1)
var cfg = config.Config{
	Client:         pokeClient,
	CaughtPokemons: map[string]pokeapi.Pokemon{},
}

func main() {
	StartCLI(&cfg)
}
