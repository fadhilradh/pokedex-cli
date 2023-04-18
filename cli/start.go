package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Start() {

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
