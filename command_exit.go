package main

import (
	"fmt"
	"os"
)

// Defining soul of ecit command
func commandExit(cfg *config, name ...string) error {
	fmt.Println()
	fmt.Println("Visit again to know more about pokemons")
	os.Exit(0)
	return nil
}
