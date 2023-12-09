package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// requests PokeAPI and gets a list of pokemon in a specific location
func (c *Client) ListLocationAreasPokemon(pageURL *string) (LocationAreasPokemonResp, error) {
	fullURL := *pageURL

	if val, ok := c.cache.Get(fullURL); ok {
		locationPokemonResp := LocationAreasPokemonResp{}
		err := json.Unmarshal(val, &locationPokemonResp)
		if err != nil {
			return LocationAreasPokemonResp{}, err
		}
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasPokemonResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasPokemonResp{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasPokemonResp{}, err
	}

	locationPokemonResp := LocationAreasPokemonResp{}
	err = json.Unmarshal(dat, &locationPokemonResp)
	if err != nil {
		return LocationAreasPokemonResp{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationPokemonResp, nil
}
