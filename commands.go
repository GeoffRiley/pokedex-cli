package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func commandExit(cfg *Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range commandIndex {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandCfg(cfg *Config, args ...string) error {
	fmt.Printf("%#v\n", cfg.NextURL)
	fmt.Printf("%#v\n", cfg.PreviousURL)
	fmt.Printf("%#v\n", cfg.pokeapiClient.CacheSize())
	return nil
}

func commandMap(cfg *Config, args ...string) error {
	var url *string
	var position string
	if args[0] == "map" {
		url = cfg.NextURL
		position = "last"
	} else {
		url = cfg.PreviousURL
		position = "first"
	}

	if url == nil {
		return fmt.Errorf("you're on the %s page", position)
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(url)
	if err != nil {
		return err
	}

	cfg.NextURL = locationsResp.Next
	cfg.PreviousURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(cfg *Config, args ...string) error {

	if len(args) < 2 {
		return errors.New("must supply an area name or number")
	}

	area := args[1]
	locationResp, err := cfg.pokeapiClient.ListLocationArea(area)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationResp.Location.Name)
	locLen := len(locationResp.PokemonEncounters)
	if locLen == 0 {
		fmt.Println("No Pokemon found here")
	}
	for i := 0; i < locLen; i++ {
		fmt.Println("-", locationResp.PokemonEncounters[i].Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *Config, args ...string) error {
	if len(args) < 2 {
		return errors.New("must supply a pokemon name")
	}
	pokeName := args[1]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokeName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokeName)

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("Failed to catch %s.\n", pokeName)
	}
	cfg.caughtPokemon[pokeName] = pokemon
	fmt.Printf("Caught %s!\n", pokeName)
	return nil
}

func commandListCaught(cfg *Config, args ...string) error {
	for name, pokemon := range cfg.caughtPokemon {
		fmt.Printf("%s: %s\n", name, pokemon.Name)
	}
	return nil
}

func commandInspect(cfg *Config, args ...string) error {
	if len(args) < 2 {
		return errors.New("must supply a pokemon name")
	}
	pokeName := args[1]
	pokemon, ok := cfg.caughtPokemon[pokeName]
	if !ok {
		return fmt.Errorf("%s is not caught", pokeName)
	}
	//fmt.Printf("%#v\n", pokemon)
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Print("Types: ")
	for _, typ := range pokemon.Types {
		fmt.Printf(" %s", typ.Type.Name)
	}
	fmt.Println()
	return nil
}
