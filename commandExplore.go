package main

import (
	"errors"
	"fmt"
)

func CommandExplore(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("explore command needs 1 argument")
	}
	fmt.Printf("Exploring %s...\n", args[0])
	locationInfo, err := config.client.GetLocation(args[0])
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range locationInfo.PokemonEncounters {
		fmt.Println("- " + encounter.Pokemon.Name)
	}
	return nil
}
