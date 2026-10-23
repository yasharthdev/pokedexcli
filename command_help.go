package main

import "fmt"

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")		
	fmt.Println("Usage:\n")
	// now I can access the registry using c.commands
	for cmd, info := range c.commands {
		fmt.Printf("%s: %s\n", cmd, info.description)
	}
	return nil
}