package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, name ...string) error {
	fmt.Println()
	fmt.Println("Visit again to know more about pokemons")
	os.Exit(0)
	return nil
}
