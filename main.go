package main

import (
	"time"

	"github.com/fadhilradh/pokedex-cli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := config{
		Client: pokeClient,
	}
	Start(&cfg)
}
