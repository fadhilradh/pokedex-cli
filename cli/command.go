package cli

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func GetCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			// callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			// callback:    commandExit,
		},
	}

}
