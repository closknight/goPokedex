package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/closknight/goPokedex/internal/pokeapi"
)

type Pokemon struct {
	name string
	info pokeapi.PokemonResponse
}

func CommandCatch(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("catch needs a pokemon name")
	}
	pokeInfo, err := config.client.GetPokemon(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s\n", pokeInfo.Name)
	if rand.Intn(pokeInfo.BaseExperience) > 30 {
		fmt.Printf("%s was caught!\n", pokeInfo.Name)
		config.pokemon[pokeInfo.Name] = Pokemon{
			name: pokeInfo.Name,
			info: pokeInfo,
		}
	} else {
		fmt.Printf("%s escaped!\n", pokeInfo.Name)
	}
	return nil
}
