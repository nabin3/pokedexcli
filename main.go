package main

import (
	"time"

	"github.com/nabin3/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.RespPokemon),
	}

	startRepl(cfg)
}
