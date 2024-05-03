package main

// Struct for representing each command
type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

// This function will return map,holds commands of pokdexcli
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
			name:        "map",
			description: "Print 20 location areas of pokemon world",
			callback:    commandMapf,
		},

		"mapb": {
			name:        "mapb",
			description: "Print previous 20 location areas of pokemon world",
			callback:    commandMapb,
		},

		"explore": {
			name:        "explore",
			description: "Print pokemons who are available in a given area",
			callback:    commandExplore,
		},

		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon",
			callback:    commandCatch,
		},

		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon",
			callback:    commandInspect,
		},

		"pokedex": {
			name:        "pokedex",
			description: "List all caught pokemons",
			callback:    commandPokedex,
		},
	}
}
