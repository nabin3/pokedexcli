package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + name

	// If url's response present in cache then get the the response from there
	if val, ok := c.cache.Get(url); ok {
		pokemonResp := RespPokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return RespPokemon{}, err
		}

		return pokemonResp, nil
	}

	// Creating an instance of http.Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	// Requesting
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer resp.Body.Close()

	// Trying to read the json from response
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	pokemonResp := RespPokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespPokemon{}, err
	}

	// Adding the new json in cache
	c.cache.Add(url, dat)

	return pokemonResp, nil
}
