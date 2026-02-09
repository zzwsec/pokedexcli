package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zzwsec/pokedexcli/internal/pokeapi"
)

type config struct {
	nextURL       *string
	previousURL   *string
	args          []string
	pokeapiClient pokeapi.Client
	pokedex       pokeapi.Pokedex
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

var cmdList map[string]cliCommand

func startRepl(cfg *config) {
	cmdList = make(map[string]cliCommand, 0)
	addCommand("exit", "Exit the Pokedex", commandExit)
	addCommand("help", "Displays a help message", commandHelp)
	addCommand("map", "Displays the next 20 locations", commandMap)
	addCommand("mapb", "Displays the prev 20 locations", commandMapB)
	addCommand("explore", "Explore a location", commandExplore)
	addCommand("catch", "Try catch pokemon", commandCatch)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")

	for scanner.Scan() {
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			fmt.Print("Pokedex > ")
			continue
		}
		commandName := words[0]
		cfg.args = words[1:]
		cmd, exists := cmdList[commandName]
		if !exists {
			fmt.Println("Unknown command")
			fmt.Print("Pokedex > ")
			continue
		}
		err := cmd.callback(cfg)
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Print("Pokedex > ")
	}

	fmt.Println()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}

func cleanInput(text string) []string {
	out := strings.ToLower(text)
	words := strings.Fields(out)
	return words
}

func addCommand(name, desc string, callback func(*config) error) {
	cmdList[name] = cliCommand{
		name:        name,
		description: desc,
		callback:    callback,
	}
}
