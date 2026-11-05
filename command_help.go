package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("map: Display the next 20 location areas")
	fmt.Println("mapb: Display the previous 20 location areas")
	fmt.Println("explore <area_name>: Display the Pokemon in a location area")
	fmt.Println("catch <pokemon>: Try to catch a Pokemon")
	fmt.Println("inspect <pokemon>: Display details about a caught Pokemon")
	fmt.Println("pokedex: List the Pokemon you have caught")
	fmt.Println("exit: Exit the Pokedex")

	return nil
}
