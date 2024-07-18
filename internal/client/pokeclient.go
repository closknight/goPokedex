package client

import (
	"encoding/json"
	"time"

	"github.com/closknight/goPokedex/internal/pokeapi"
	"github.com/closknight/goPokedex/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache pokecache.Cache
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
	}
}

func (c Client) GetLocations(url *string) (pokeapi.LocationsResponse, error) {
	if url == nil {
		url = new(string)
		*url = baseURL + "/location-area/"
	}
	if data, ok := c.cache.Get(*url); ok {
		locationsRes := pokeapi.LocationsResponse{}
		err := json.Unmarshal(data, &locationsRes)
		if err != nil {
			return pokeapi.LocationsResponse{}, err
		}
		return locationsRes, nil
	}

	resp, err := pokeapi.GetLocations(*url)
	if err != nil {
		return pokeapi.LocationsResponse{}, err
	}
	data, err := json.Marshal(resp)
	if err != nil {
		return pokeapi.LocationsResponse{}, err
	}
	c.cache.Add(*url, data)
	return resp, nil
}

func (c Client) GetLocation(name string) (pokeapi.ExploreResponse, error) {
	url := baseURL + "/location-area/" + name

	if data, ok := c.cache.Get(url); ok {
		exploreResp := pokeapi.ExploreResponse{}
		err := json.Unmarshal(data, &exploreResp)
		if err != nil {
			return pokeapi.ExploreResponse{}, err
		}
		return exploreResp, nil
	}

	exploreResponse, err := pokeapi.GetExploreLocation(url)
	if err != nil {
		return pokeapi.ExploreResponse{}, err
	}
	data, err := json.Marshal(exploreResponse)
	if err != nil {
		return pokeapi.ExploreResponse{}, err
	}
	c.cache.Add(url, data)
	return exploreResponse, nil
}
