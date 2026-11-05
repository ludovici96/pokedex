package main

import (
	"fmt"
)

type locationAreaDetailResponse struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: explore <area_name>")
	}

	name := args[0]

	fmt.Printf("Exploring %s...\n", name)

	data, err := fetch[locationAreaDetailResponse](cfg, locationAreaURL+"/"+name)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range data.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
