package main

import (
	"errors"
	"fmt"
)

func CommandInspect(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("inspect needs 1 pokemon")
	}

	pokemonInfo, ok := config.pokemon[args[0]]

	if !ok {
		return errors.New(args[0] + " has not been caught")
	}
	fmt.Printf("Name: %s\n", pokemonInfo.info.Name)
	fmt.Printf("Height: %v\n", pokemonInfo.info.Height)
	fmt.Printf("Weight: %v\n", pokemonInfo.info.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonInfo.info.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pType := range pokemonInfo.info.Types {
		fmt.Printf("  - %s\n", pType.Type.Name)
	}
	return nil
}
