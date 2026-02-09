package main

import (
	"time"

	"github.com/zzwsec/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Second*5, time.Minute*3)
	pokedex := pokeapi.NewPokedex()
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       pokedex,
	}
	startRepl(cfg)
}
