package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreasPokemon(pageURL *string) (LocationAreasPokemonResp, error) {
	fullURL := *pageURL

	//check if in cache

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

	//add to cache

	return locationPokemonResp, nil
}
