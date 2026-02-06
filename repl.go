package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/zzwsec/pokedexcli/internal/pokecache"
)

type config struct {
	nextURL     *string
	previousURL *string
	cache       pokecache.Cache
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

var cmdList map[string]cliCommand

func startRepl() {
	cmdList = make(map[string]cliCommand, 0)
	cfg := &config{
		cache: pokecache.NewCache(time.Second * 5),
	}
	addCommand("exit", "Exit the Pokedex", commandExit)
	addCommand("help", "Displays a help message", commandHelp)
	addCommand("map", "Displays the next 20 locations", commandMap)
	addCommand("mapb", "Displays the prev 20 locations", commandMapB)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")

	for scanner.Scan() {
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			fmt.Print("Pokedex > ")
			continue
		}
		commandName := words[0]
		switch commandName {
		case "exit":
			_ = cmdList["exit"].callback(cfg)
		case "help":
			_ = cmdList["help"].callback(cfg)
		case "map":
			_ = cmdList["map"].callback(cfg)
		case "mapb":
			_ = cmdList["mapb"].callback(cfg)
		default:
			fmt.Println("Unknown command")
		}
		fmt.Print("Pokedex > ")
	}

	fmt.Println()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}

func cleanInput(text string) []string {
	strSlice := strings.Fields(text)
	for i, str := range strSlice {
		strSlice[i] = strings.ToLower(str)
	}
	return strSlice
}

func addCommand(name, desc string, callback func(*config) error) {
	cmdList[name] = cliCommand{
		name:        name,
		description: desc,
		callback:    callback,
	}
}
