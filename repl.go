package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var cmdList map[string]cliCommand

func startRepl() {
	cmdList = make(map[string]cliCommand, 0)
	addCommand("exit", "Exit the Pokedex", commandExit)
	addCommand("help", "Displays a help message", commandHelp)

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
			_ = cmdList["exit"].callback()
		case "help":
			_ = cmdList["help"].callback()
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

func addCommand(name, desc string, callback func() error) {
	cmdList[name] = cliCommand{
		name:        name,
		description: desc,
		callback:    callback,
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage")
	fmt.Println()
	for _, cmd := range cmdList {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
