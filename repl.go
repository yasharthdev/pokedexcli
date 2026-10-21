package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func cleanInput(text string) []string {
	// make the text lowercase first and then split the string
	return strings.Fields(strings.ToLower(text))
}

func startRepl() {
	// read from Stdin
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			// this is a boolean, which if false means eof or
			//perhaps an error occurred
			break
		}	
		
		user_input := cleanInput(scanner.Text())
		if len(user_input) == 0 {
			// if user presses enter, the program doesn't panic
			continue
		}
		
		fmt.Printf("Your command was: %s\n", user_input[0])
	}
	
	// check for errors after breaking out
	if err := scanner.Err(); err != nil {
		fmt.Fprint(os.Stderr, "Error reading input: ", err)
	}
}

