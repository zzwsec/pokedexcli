package main

import (
	"fmt"
	"os"

	"github.com/zzwsec/pokedexcli/internal/pokeapi"
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
	resp, err := pokeapi.GetLocationAreas(cfg.nextURL, &cfg.cache)
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
	// 检查是否在第一页
	if cfg.previousURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	resp, err := pokeapi.GetLocationAreas(cfg.previousURL, &cfg.cache)
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
