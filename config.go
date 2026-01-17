package main

import "github.com/GeoffRiley/pokedex-cli/internal/pokeapi"

type Config struct {
	NextURL       *string
	PreviousURL   *string
	pokeapiClient pokeapi.Client
	caughtPokemon map[string]pokeapi.Pokemon
}

type Result_arr struct {
	Name string `json:name`
	Url  string `json:url`
}

type Response struct {
	Count    int          `json:count`
	Next     string       `json:next`
	Previous string       `json:prev`
	Results  []Result_arr `json:results`
}
