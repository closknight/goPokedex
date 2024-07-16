package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/closknight/goPokedex/internal/pokeapi"
)

func CommandMap(config *config) error {
	if config.next == nil && config.prev != nil {
		return errors.New("no more locations")
	}
	if config.next == nil {
		config.next = new(string)
		*config.next = pokeapi.GetLocationsURL()
	}

	data, isStored := config.cache.Get(*config.next)
	var locationsRes pokeapi.LocationsResponse
	if isStored {
		err := json.Unmarshal(data, &locationsRes)
		if err != nil {
			return nil
		}
	} else {
		resp, err := pokeapi.GetLocations(config.next)
		if err != nil {
			return err
		}
		data, err := json.Marshal(resp)
		if err != nil {
			return err
		}
		config.cache.Add(*config.next, data)
		locationsRes = resp
	}
	config.next = locationsRes.Next
	config.prev = locationsRes.Previous

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func CommandMapb(config *config) error {
	if config.prev == nil {
		return errors.New("no previous locations")
	}

	var locationsRes pokeapi.LocationsResponse
	data, isStored := config.cache.Get(*config.prev)
	if isStored {
		err := json.Unmarshal(data, &locationsRes)
		if err != nil {
			return err
		}
	} else {
		resp, err := pokeapi.GetLocations(config.prev)
		if err != nil {
			return err
		}
		data, err := json.Marshal(resp)
		if err != nil {
			return err
		}
		config.cache.Add(*config.prev, data)
		locationsRes = resp
	}

	config.next = locationsRes.Next
	config.prev = locationsRes.Previous

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
