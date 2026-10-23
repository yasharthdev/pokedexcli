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

func startRepl(conf *config) {
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
		
		// conf.commands gives me the registry, gotta check if the
		// user typed in a valid command
		desc, ok := conf.commands[user_input[0]]
		if ok {
			if err := desc.callback(conf); err != nil {
				fmt.Printf("Failed to call the callback func: %v\n", err)	
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
	callback func(*config) error
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
		"map": {
			name: "map command",
			description: "Display a batch of location areas on the map",
			callback: commandMap,
		},
		"mapb": {
			name: "map back command",
			description: "Display the previous 20 location areas on the map",
			callback: commandMapBack,
		},
	}
	
	return commandRegistry
}

type config struct {
	commands map[string]cliCommand
	nextURL *string
	previousURL *string
}

