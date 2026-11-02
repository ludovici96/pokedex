package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const locationAreaURL = "https://pokeapi.co/api/v2/location-area"

type locationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg *config) error {
	url := locationAreaURL

	if cfg.nextLocationsURL != nil {
		url = *cfg.nextLocationsURL
	}

	body, ok := cfg.pokeapiCache.Get(url)

	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("API request failed: %s", res.Status)
		}

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cfg.pokeapiCache.Add(url, body)
	}

	var data locationAreaResponse

	err := json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = data.Next
	cfg.prevLocationsURL = data.Previous

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}

	return nil
}
