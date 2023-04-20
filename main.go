package main

import (
	"time"

	"github.com/fadhilradh/pokedex-cli/internal/pokeapi"
)

var pokeClient = pokeapi.NewClient(5*time.Second, time.Minute*1)
var cfg = config{
	Client: pokeClient,
}

func main() {
	StartCLI(&cfg)
}
