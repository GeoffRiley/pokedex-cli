package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commandIndex map[string]cliCommand

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		err := scanner.Scan()
		if !err {
			fmt.Println("Done.")
			return
		}
		inp := scanner.Text()
		words := cleanInput(inp)
		command := words[0]
		ctrl, ok := commandIndex[command]
		if ok {
			_ = ctrl.callback()
		} else {
			fmt.Println("Unknown command")
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
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}
