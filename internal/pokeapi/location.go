package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"time"
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

func (c *Client) CatchPokemon(pkg *Pokedex, baseUrl, params string) (bool, error) {
	reqUrl := "https://pokeapi.co/api/v2/pokemon/"
	if baseUrl != "" {
		reqUrl = baseUrl
	}
	u, err := url.JoinPath(reqUrl, params)
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return false, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, err
	}

	p := Pokemon{}
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return false, err
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	exp := p.BaseExperience
	if exp < 1 {
		exp = 1
	}
	roll := r.Intn(exp + 1)
	if roll < 40 {
		pkg.mu.Lock()
		pkg.pkg[p.Name] = p
		pkg.mu.Unlock()
		return true, nil
	}
	return false, nil
}
