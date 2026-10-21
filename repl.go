package main

import "strings"

func cleanInput(text string) []string {
	// make the text lowercase first and then split the string
	return strings.Fields(strings.ToLower(text))
}

