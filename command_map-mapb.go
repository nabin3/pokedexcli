package main

import (
	"fmt"

	"github.com/nabin3/pokedexcli/apinTeraction"
)

// Creating pointer of an instance of ConfigUrls struct, which will be reciever for FetchLocationAreas from apinTeraction package
var navigator *apinTeraction.ConfigUrls = &apinTeraction.ConfigUrls{
	Next:     "https://pokeapi.co/api/v2/location-area/",
	Previous: "",
}

func commandMap() error {

	// If reached end page of location-areas
	if navigator.Next == "" {
		fmt.Println("Can't proceed further")
		return nil
	}

	// Making call to location-area endpoint of pokedex api with help of FetchLocationAreas func
	nextList, err := navigator.FetchLocationAreas(true)

	// Checking ig any error occured during the api call
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	// Printing each location-area name
	for _, name := range nextList {
		fmt.Println(name)
	}

	return nil
}

func commandMapb() error {

	// If reached start page of location-areas
	if navigator.Previous == "" {
		fmt.Println("Can't get back further")
		return nil
	}

	// Making call to location-area endpoint of pokedex api with help of FetchLocationAreas func
	previousList, err := navigator.FetchLocationAreas(false)

	// Checking ig any error occured during the api call
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	// Printing each location-area name
	for _, name := range previousList {
		fmt.Println(name)
	}

	return nil
}
