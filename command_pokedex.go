package main

import "fmt"

// Defining soul of pokedex command
func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex")

	// Iterating caughtPokemon map and print every pokemon's name here
	for name := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
