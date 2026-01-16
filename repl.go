package main

import "strings"

func cleanInput(text string) []string {
	var words []string
	fields := strings.Fields(text)
	for _, word := range fields {
		words = append(words, strings.ToLower(word))
	}
	return words
}


