package main

import (
	"fmt"
	"math/rand"
)

const pokemonURL = "https://pokeapi.co/api/v2/pokemon"

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: catch <pokemon>")
	}

	name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := fetch[Pokemon](cfg, pokemonURL+"/"+name)
	if err != nil {
		return err
	}

	if rand.Intn(pokemon.BaseExperience+1) >= 50 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	cfg.pokedex[pokemon.Name] = pokemon

	return nil
}
