package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetLocationAreas(pageURL *string) (*LocationAreasResponse, error) {
	url := "https://pokeapi.co/api/v2/location-area?limit=20&offset=0"
	if pageURL != nil {
		url = *pageURL
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	las := &LocationAreasResponse{}
	err = json.NewDecoder(resp.Body).Decode(las)
	if err != nil {
		return nil, fmt.Errorf("JSON parsing failed: %w", err)
	}
	return las, nil
}
