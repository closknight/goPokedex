package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2"

type LocationsResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationsURL() string {
	return baseURL + "/location-area"
}

func GetLocations(webpage *string) (LocationsResponse, error) {
	url := baseURL + "/location-area"
	if webpage != nil {
		url = *webpage
	}
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
