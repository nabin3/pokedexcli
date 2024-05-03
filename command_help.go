package main

import (
	"fmt"
)

// Defining help command's soul
func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Print("Availablke Pokedex! commands:\n")
	fmt.Println()
	// Iterate on each command stored on map which hold commands
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s \n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
