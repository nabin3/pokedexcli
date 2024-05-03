package main

import (
	"errors"
	"fmt"
	"strings"
)

// Defining soul of explore command
func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("please give an area-name")
	}

	// Obtaining area name
	areaName := args[0]

	// retrieveing all pkemons of a given area
	pokemonsResp, err := cfg.pokeapiClient.ListPokemons(areaName)
	if err != nil {
		// This error was appearing on invalid name, so this is a quick fix
		if strings.Contains(err.Error(), "invalid character 'N' looking for beginning of value") {
			return errors.New("try again with a valid area name")
		}
		return err
	}

	fmt.Printf("Searching in %s .... \n Pokemon founds \n", areaName)

	// Listing the pokemons found
	for _, item := range pokemonsResp.PokemonEncounters {
		fmt.Printf("--%s \n", item.Pokemon.Name)
	}

	return nil
}
