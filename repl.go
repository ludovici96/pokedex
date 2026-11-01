package main

import (
	"strings"
	"fmt"
	"os"
	"bufio"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}

		firstWord := words[0]

		fmt.Printf("Your command was: %s\n", firstWord)
	}

}

func cleanInput(text string) []string  {
	cleaned := strings.Fields(strings.ToLower(strings.TrimSpace(text)))
	return cleaned
}