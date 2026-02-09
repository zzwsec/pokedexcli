package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range cmdList {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextURL)
	if err != nil {
		return err
	}
	cfg.nextURL = resp.Next
	cfg.previousURL = resp.Previous
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapB(cfg *config) error {
	if cfg.previousURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.previousURL)
	if err != nil {
		return err
	}
	cfg.nextURL = resp.Next
	cfg.previousURL = resp.Previous
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandExplore(cfg *config) error {
	orgUrl := "https://pokeapi.co/api/v2/location-area"

	if len(cfg.args) != 1 {
		return fmt.Errorf("Please provide a location area name or id")
	}
	resp, err := cfg.pokeapiClient.GetLocationPokemon(&orgUrl, &cfg.args[0])
	if err != nil {
		return err
	}
	if resp.Name == "" {
		fmt.Printf("Invalid location\n")
		return nil
	}
	fmt.Printf("Exploring %s...\n", resp.Name)
	fmt.Printf("Found Pokemon:\n")
	for _, pe := range resp.PokemonEncounters {
		fmt.Printf("  - %s\n", pe.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config) error {
	orgUrl := "https://pokeapi.co/api/v2/pokemon/"

	if len(cfg.args) != 1 {
		return fmt.Errorf("Please provide a pokemon name or id")
	}

	fmt.Printf("Throwing a Pokeball at %s\n", cfg.args[0])
	ok, err := cfg.pokeapiClient.CatchPokemon(&cfg.pokedex, orgUrl, cfg.args[0])
	if err != nil {
		return err
	}

	if !ok {
		fmt.Printf("%s escaped!\n", cfg.args[0])
		return nil
	}

	fmt.Printf("%s was caught!\n", cfg.args[0])
	return nil
}
