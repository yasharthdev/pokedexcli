
package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")		
	fmt.Println("Usage:\n")
	registry := getCommands()
	for cmd, info := range registry {
		fmt.Printf("%s: %s\n", cmd, info.description)
	}
	return nil
}