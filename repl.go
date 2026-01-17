package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commandIndex map[string]cliCommand

func startRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		err := scanner.Scan()
		if !err {
			fmt.Println("Done.")
			return
		}

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		command := words[0]
		ctrl, ok := commandIndex[command]
		if ok {
			err := ctrl.callback(cfg, words...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	var words []string
	fields := strings.Fields(text)
	for _, word := range fields {
		words = append(words, strings.ToLower(word))
	}
	return words
}

func init() {
	commandIndex = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays a list of the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display a list of the previous 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"explore": {
			name:        "explore",
			description: "Explore a given area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"list": {
			name:        "list",
			description: "List caught pokemon",
			callback:    commandListCaught,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon",
			callback:    commandInspect,
		},
		"cfg": {
			name:        "cfg",
			description: "Show config",
			callback:    commandCfg,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}
