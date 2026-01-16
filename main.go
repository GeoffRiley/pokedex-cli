package main

import (
	"time"

	"github.com/GeoffRiley/pokedex-cli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &Config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
