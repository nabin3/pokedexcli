package main


type cliCommand struct {
	name        string
	description string
	callback    func() error
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},

		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},

		"map": {
			name: "map",
			description: "Print 20 location areas of pokemon world",
			callback: commandMap,
		},

		"mapb": {
			name: "mapb",
			description: "Print previous 20 location areas of pokemon world",
			callback: commandMapb,
		},		
	}
}
