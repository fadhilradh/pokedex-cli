package main

import (
	"github.com/fadhilradh/pokedex-cli/cli"
	"github.com/fadhilradh/pokedex-cli/pokedex"
)

func main() {

	cli.Start(pokedex.Cfg)

}
