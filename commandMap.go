package main

import (
	"errors"
	"fmt"
)

func CommandMap(config *config, args ...string) error {
	if config.next == nil && config.prev != nil {
		return errors.New("no more locations")
	}

	locationsRes, err := config.client.GetLocations(config.next)
	if err != nil {
		return err
	}

	config.next = locationsRes.Next
	config.prev = locationsRes.Previous

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func CommandMapb(config *config, args ...string) error {
	if config.prev == nil {
		return errors.New("no previous locations")
	}

	locationsRes, err := config.client.GetLocations(config.prev)
	if err != nil {
		return err
	}

	config.next = locationsRes.Next
	config.prev = locationsRes.Previous

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
