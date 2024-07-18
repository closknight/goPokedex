package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2"

func GetLocationsURL() string {
	return baseURL + "/location-area"
}

func GetLocations(url string) (LocationsResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationsResponse{}, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return LocationsResponse{}, err
	}
	if res.StatusCode > 299 {
		return LocationsResponse{}, fmt.Errorf("network error:\n status code: %v\n body- %s", res.StatusCode, string(body))
	}

	responseJson := LocationsResponse{}
	err = json.Unmarshal(body, &responseJson)
	if err != nil {
		return LocationsResponse{}, err
	}
	return responseJson, nil
}

func GetExploreLocation(url string) (ExploreResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return ExploreResponse{}, err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		return ExploreResponse{}, err
	}
	if res.StatusCode > 299 {
		return ExploreResponse{}, fmt.Errorf("network error:\n status code: %v\n body- %s", res.StatusCode, string(body))
	}
	exploreResp := ExploreResponse{}
	err = json.Unmarshal(body, &exploreResp)
	if err != nil {
		return ExploreResponse{}, err
	}
	return exploreResp, nil
}
