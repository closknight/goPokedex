package main

import "fmt"

func CommandPokedex(config *config, args ...string) error {
	if len(config.pokemon) == 0 {
		fmt.Println("No pokemon have been caught")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.pokemon {
		fmt.Printf("  - %s\n", pokemon.name)
	}
	return nil
}
