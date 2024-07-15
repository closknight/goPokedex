package main

import (
	"errors"
	"fmt"
)

const LOCATION_WEBPAGE = "https://pokeapi.co/api/v2/location"

func CommandMap(config *config) error {
	if config.next == nil && config.prev != nil {
		return errors.New("no more locations")
	}

	webpage := LOCATION_WEBPAGE
	if config.next != nil {
		webpage = *config.next
	}

	locationsRes, err := getLocations(webpage)
	if err != nil {
		return err
	}

	config.next = locationsRes.Next
	config.prev = locationsRes.Previous

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}
	fmt.Println()
	return nil
}

func CommandMapb(config *config) error {
	if config.prev == nil {
		return errors.New("no previous locations")
	}
	locationsRes, err := getLocations(*config.prev)
	if err != nil {
		return err
	}

	config.next = locationsRes.Next
	config.prev = locationsRes.Previous

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}
	fmt.Println()
	return nil
}
