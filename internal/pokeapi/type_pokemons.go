package pokeapi

type RespShallowPokemons struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

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