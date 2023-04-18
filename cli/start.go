package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/fadhilradh/pokedex-cli/pokedex"
)

func Start(config *pokedex.Config) {

	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		commands := GetCommand()
		commands[scanner.Text()].callback()
	}
}
