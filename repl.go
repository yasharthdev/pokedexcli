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
		
		registry := getCommands()
		// if the user types in a valid command that is present in the registry
		desc, ok := registry[user_input[0]]
		if ok {
			if err := desc.callback(); err != nil {
				fmt.Errorf("Failed to call the callback function: %v", err)
			}
		} else {
			fmt.Println("Unknown command")
		}	

	}
	
	// check for errors after breaking out
	if err := scanner.Err(); err != nil {
		fmt.Fprint(os.Stderr, "Error reading input: ", err)
	}
}

type cliCommand struct {
	name string
	description string
	callback func() error
}


func getCommands() map[string]cliCommand {
	commandRegistry := map[string]cliCommand{
		"exit": {
			name: "exit command",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help command",
			description: "How to use Pokedex",
			callback: commandHelp,
		},
	}
	
	return commandRegistry
}
