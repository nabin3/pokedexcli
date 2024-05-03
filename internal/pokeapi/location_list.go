package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area" // Constructing url
	if pageURL != nil {
		url = *pageURL // If a URL is passed to ListLocations func then url is updated
	}

	// If url's response present in cache then get the the response from there
	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsResp) // Decoding retrieved val(data) from cache
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationsResp, nil
	}

	// Creating an instance of http.Request when request URL response couldn't be found in cache
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// Requesting
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	// Trying to read the json from response
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// Decoding data retrieved from response obtained from pokeapi server
	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// Adding the new json in cache
	c.cache.Add(url, dat)

	return locationsResp, nil
}
