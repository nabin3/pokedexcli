package main

import (
	"errors"
	"fmt"
)

// Defining soul of map command
func commandMapf(cfg *config, args ...string) error {
	// Retrieveing location names
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	// Updating next and previous location url
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	// Printing each locatrion name obtained in response
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

// Defining soul of mapb command
func commandMapb(cfg *config, args ...string) error {
	// If user on first page of locations, then they can't go back further
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	// Retrieveing location names
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	// Updating next and previous locations
	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	// Print location names obtained in response
	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
