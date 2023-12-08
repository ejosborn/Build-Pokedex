package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// requests api
func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {

	fullURL := baseURL + "/location-area"
	if pageURL != nil {
		fullURL = *pageURL
	}

	if val, ok := c.cache.Get(fullURL); ok {
		locationsResp := LocationAreasResp{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationAreasResp{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationsResp := LocationAreasResp{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationsResp, nil
}
