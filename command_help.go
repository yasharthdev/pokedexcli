package main

import "fmt"

func commandHelp(conf *config) error {
	fmt.Println("Welcome to the Pokedex!")		
	fmt.Println("Usage:\n")
	// now I can access the registry using c.commands
	for cmd, info := range conf.commands {
		fmt.Printf("%s: %s\n", cmd, info.description)
	}
	return nil
}