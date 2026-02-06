package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/zzwsec/pokedexcli/internal/pokecache"
)

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

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

func normalizeURL(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	params := u.Query()
	u.RawQuery = params.Encode()
	return u.String(), nil
}
