package pokeapi

import (
	"net/http"
	"sync"

	"github.com/zzwsec/pokedexcli/internal/pokecache"
)

type LocationAreasResponse struct {
	Count    int                `json:"count"`
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []LocationAreaItem `json:"results"`
}

type LocationAreaItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

type AreaDetail struct {
	PokemonEncounters []PokemonDetail `json:"pokemon_encounters"`
	ID                int             `json:"id"`
	Name              string          `json:"name"`
}

type PokemonDetail struct {
	Pokemon struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon"`
}

type Pokedex struct {
	pkg map[string]Pokemon
	mu  *sync.RWMutex
}

type Pokemon struct {
	BaseExperience int    `json:"base_experience"`
	Name           string `json:"name"`
}
