package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type locationsResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getLocations(webpage string) (*locationsResponse, error) {
	fmt.Print("webpage: ")
	fmt.Println(webpage)
	res, err := http.Get(webpage)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	if res.StatusCode > 299 {
		return nil, errors.New(fmt.Sprintf("Network Error:\n Status code: %v\n body- %s", res.StatusCode, string(body)))
	}

	responseJson := locationsResponse{}
	err = json.Unmarshal(body, &responseJson)
	if err != nil {
		return nil, err
	}
	return &responseJson, nil
}
