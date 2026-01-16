package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range commandIndex {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *Config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.NextURL)
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

func commandMapb(cfg *Config) error {
	if cfg.PreviousURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.PreviousURL)
	if err != nil {
		return err
	}

	cfg.NextURL = locationResp.Next
	cfg.PreviousURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
