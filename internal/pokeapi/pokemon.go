package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) CatchPokemon(pokemonName string) (Pokemon, error) {
	u, err := url.JoinPath(baseURL, "pokemon", pokemonName)
	if err != nil {
		return Pokemon{}, err
	}

	fmt.Printf("[DEBUG] Cache key: %s\n", u)
	if cacheData, ok := c.cache.Get(u); ok {
		p := Pokemon{}
		if err := json.Unmarshal(cacheData, &p); err == nil {
			fmt.Printf("[DEBUG] ✅ CACHE HIT! Using cached data\n")
			return p, nil
		}
	}

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Pokemon{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	p := Pokemon{}
	err = json.Unmarshal(data, &p)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(u, data)
	return p, nil
}

func (p *Pokedex) Get(name string) (Pokemon, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	v, ok := p.pkg[name]
	return v, ok
}

func (p *Pokedex) Set(pp Pokemon) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.pkg[pp.Name] = pp
}

func (p *Pokedex) List() map[string]Pokemon {
	p.mu.RLock()
	defer p.mu.RUnlock()
	out := make(map[string]Pokemon, len(p.pkg))
	for k, v := range p.pkg {
		out[k] = v
	}
	return out
}
