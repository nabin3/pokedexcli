package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

// Defining soul of catch command
func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]

	// Checking if we alrady have caught the given pokemon
	if _, exist := cfg.caughtPokemon[name]; exist {
		return errors.New("given pokemon is already in your pokedex")
	}
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		// This error was appearing on invalid name, so this is a quick fix
		if strings.Contains(err.Error(), "invalid character 'N' looking for beginning of value") {
			return errors.New("try again with a valid pokemon name")
		}
		return err
	}

	// Upon base_experience of a pokemon we generate a non-neagative random number, which will determine if we can catch the given pokemon
	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Forms[0].Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Forms[0].Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemon.Forms[0].Name)

	// Storing the caught pokemon in caughtPokemon map
	cfg.caughtPokemon[pokemon.Forms[0].Name] = pokemon
	return nil
}
