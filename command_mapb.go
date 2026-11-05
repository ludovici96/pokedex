package main

import (
	"fmt"
)

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return fmt.Errorf("you're on the first page")
	}

	url := *cfg.prevLocationsURL

	data, err := fetchLocationAreas(cfg, url)
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