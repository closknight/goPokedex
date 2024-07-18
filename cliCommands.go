package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *config, args ...string) error
}

func commandHelp(config *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit(config *config, args ...string) error {
	os.Exit(0)
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "prints out helpful message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "ends program",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "gives the next 20 locations of the pokemon world",
			callback:    CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "gives the previous 20 locations of the pokemon world",
			callback:    CommandMapb,
		},
		"explore": {
			name:        "explore <area_name>",
			description: "prints out the pokemon that can be found at `area_name`",
			callback:    CommandExplore,
		},
	}
}
