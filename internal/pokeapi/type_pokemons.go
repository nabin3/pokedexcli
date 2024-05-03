package pokeapi

// BluePrint for retrieving pokemon names from pokeapi response(in JSON format)
type RespShallowPokemons struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

// BluePrint for extracting pokemon's details from pokapi response
type RespPokemon struct {
	Forms []struct {
		Name string `json:"name"`
	} `json:"forms"`

	BaseExperience int `json:"base_experience"`

	Height int `json:"height"`

	Weight int `json:"weight"`

	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`

	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}
