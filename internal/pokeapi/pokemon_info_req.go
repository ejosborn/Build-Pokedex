package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// requests PokeAPI and gets info
func (c *Client) ListPokemonInfo(pageURL *string) (PokemonInfoResp, error) {
	fullURL := *pageURL

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonInfoResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfoResp{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInfoResp{}, err
	}

	pokemonInfoResp := PokemonInfoResp{}
	err = json.Unmarshal(dat, &pokemonInfoResp)
	if err != nil {
		return PokemonInfoResp{}, err
	}

	return pokemonInfoResp, nil
}
