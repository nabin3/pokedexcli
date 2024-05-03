package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nabin3/pokedexcli/internal/pokeapi"
)

// REPL configuration struct
type config struct {
	pokeapiClient    pokeapi.Client                 // BluePrint of our client
	caughtPokemon    map[string]pokeapi.RespPokemon // In this map we will store all pokemons that have been caught
	nextLocationsURL *string                        // This represent url of api endpoint which will serve next pokemons
	prevLocationsURL *string                        // This represent url of api endpoint which will serve previous pokemons
}

// Our REPL func defination
func startRepl(cfg *config) {
	// Creating a new sacanner to read input from console
	reader := bufio.NewScanner(os.Stdin)

	// This infinte for loop is responsible to crete the REPL behaviour
	for {
		fmt.Print("Pokedex > ")

		// Advancing our reader scanner to next token
		reader.Scan()

		// Reading and cleaning input and making a string-slice which will hold each input word
		words := cleanInput(reader.Text())

		// If user has not issued any command and pressed enter tthen we do thing and advance to next iteration of our REPL
		if len(words) == 0 {
			continue
		}

		// Retrieveing command name from input line
		commandName := words[0]
		// Creating empty string slice which will be used to hold command arguments
		args := []string{}
		if len(words) > 1 { // If any argument is given args slice will store those arguments
			args = words[1:] // Extracting arguments from input
		}

		// Retrieveng map which holds commands, checking and retrieving if user entered command esists in the map or not
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...) // Executing the command

			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

// This func will clean input by convertingeach input letter in lowercase and put every word of input string as a single string in a slice of string
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
