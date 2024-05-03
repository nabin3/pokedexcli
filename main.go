package main

import (
	"fmt"
	"time"

	"github.com/nabin3/pokedexcli/internal/pokeapi"
)

func main() {
	// Creating new client with timeout of 5 seconds and cache of 5 minutes
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)

	// Configuring our REPL
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.RespPokemon),
	}

	// Welcome msg
	fmt.Println("Welcome to Pokedex,personal assistance to explore pokemon world\n\ntype help to see available commands")

	// Starting our REPL
	startRepl(cfg)
}
