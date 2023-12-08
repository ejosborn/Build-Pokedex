package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// requests api
func (c *Client) ListLocationAreas(pageURL *string) (RespShallowLocations, error) {

	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	if val, ok := c.cache.get(fullURL); ok {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return RespShallowLocations{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.add(fullURL, dat)

	return locationsResp, nil

}
