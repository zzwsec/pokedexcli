package pokeapi

import (
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

func normalizeURL(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	params := u.Query()
	u.RawQuery = params.Encode()
	return u.String(), nil
}
