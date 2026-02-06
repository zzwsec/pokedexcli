package main

import (
	"time"

	"github.com/zzwsec/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Second*5, time.Minute*3)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
