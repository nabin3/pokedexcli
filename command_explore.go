package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("please give an area-name")
	}

	areaName := args[0]

	pokemonsResp, err := cfg.pokeapiClient.ListPokemons(areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Searching in %s \n Pokemon founds \n", areaName)
	for _, item := range pokemonsResp.PokemonEncounters {
		fmt.Printf("--%s \n", item.Pokemon.Name)
	}

	return nil
}
