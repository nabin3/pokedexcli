package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListPokemons -
func (c *Client) ListPokemons(area string) (RespShallowPokemons, error) {
	url := baseURL + "/location-area/" + area // Constructing URL

	// If url's response present in cache then get the the response from there
	if val, ok := c.cache.Get(url); ok {
		pokemonsResp := RespShallowPokemons{}
		err := json.Unmarshal(val, &pokemonsResp)
		if err != nil {
			return RespShallowPokemons{}, err
		}

		return pokemonsResp, nil
	}

	// Creating an instance of http.Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemons{}, err
	}

	// Requesting
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemons{}, err
	}
	defer resp.Body.Close()

	// Trying to read the json from response
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowPokemons{}, err
	}

	pokemonsResp := RespShallowPokemons{}
	err = json.Unmarshal(dat, &pokemonsResp)
	if err != nil {
		return RespShallowPokemons{}, err
	}

	// Adding the new json in cache
	c.cache.Add(url, dat)

	return pokemonsResp, nil
}
