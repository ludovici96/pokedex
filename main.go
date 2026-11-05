package main

import (
	"time"
	"github.com/ludovici96/pokedex/internal/pokecache"
)

func main() {
	cfg := &config{
		commands: getCommands(),
		pokeapiCache: pokecache.NewCache(5 * time.Minute),
		pokedex: make(map[string]Pokemon),
	}

	startRepl(cfg)
}
