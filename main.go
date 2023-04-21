package main

import (
	"time"

	"github.com/fadhilradh/pokedex-cli/internal/pokeapi"
)

var pokeClient = pokeapi.NewClient(10*time.Second, time.Hour*1)
var cfg = config{
	Client: pokeClient,
}

func main() {
	StartCLI(&cfg)
}
