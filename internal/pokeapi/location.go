package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	url := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}
	cacheKey, err := normalizeURL(url)
	if err != nil {
		cacheKey = url
	}

	if cacheData, ok := c.cache.Get(cacheKey); ok {
		las := LocationAreasResponse{}
		if err := json.Unmarshal(cacheData, &las); err == nil {
			return las, nil
		}
	}

	req, err := http.NewRequest("GET", cacheKey, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationAreasResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	las := LocationAreasResponse{}
	err = json.Unmarshal(data, &las)
	if err != nil {
		return LocationAreasResponse{}, fmt.Errorf("JSON parsing failed: %w", err)
	}

	c.cache.Add(cacheKey, data)
	return las, nil
}

func (c *Client) GetLocationPokemon(baseUrl, params *string) (AreaDetail, error) {
	orgUrl := "https://pokeapi.co/api/v2/location-area"
	if baseUrl != nil {
		orgUrl = *baseUrl
	}

	fixUrl, _ := url.JoinPath(orgUrl, *params)
	fmt.Printf("[DEBUG] Cache key: %s\n", fixUrl)
	if cacheData, ok := c.cache.Get(fixUrl); ok {
		ad := AreaDetail{}
		if err := json.Unmarshal(cacheData, &ad); err == nil {
			fmt.Printf("[DEBUG] ✅ CACHE HIT! Using cached data\n")
			return ad, nil
		}
	}

	fmt.Printf("[DEBUG] ❌ CACHE MISS - Fetching from API...\n")
	req, err := http.NewRequest("GET", fixUrl, nil)
	if err != nil {
		return AreaDetail{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AreaDetail{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return AreaDetail{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return AreaDetail{}, err
	}
	ad := AreaDetail{}
	err = json.Unmarshal(data, &ad)
	if err != nil {
		return AreaDetail{}, err
	}

	c.cache.Add(fixUrl, data)
	return ad, nil
}
