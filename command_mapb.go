package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return fmt.Errorf("you're on the first page")
	}

	url := *cfg.prevLocationsURL

	// Check cache first
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

	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	cfg.nextLocationsURL = data.Next
	cfg.prevLocationsURL = data.Previous

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}

	return nil
}