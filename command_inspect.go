package main

import (
	"errors"
	"fmt"
)

// Defining soul of inspect command
func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	// Obtaining desired pokemon from caughtPokemon map
	name := args[0]
	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	// Printing information about obatined pokemon
	fmt.Println("Name:", pokemon.Forms[0].Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats { // printing stats of obtained pokemon
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types { // Printing types of obtained pokemon
		fmt.Println("  -", typeInfo.Type.Name)
	}
	return nil
}
