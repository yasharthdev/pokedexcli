package main

import (
	"fmt"
	"github.com/yasharthdev/pokedexcli/internal/pokeapi"
	"log"
)

func commandMapBack(conf *config) error {
	if conf.previousURL == nil {
		fmt.Println("you're on the first page")
		return nil
	} 
	locationArea, err := pokeapi.GetLocationAreas(conf.previousURL)
	if err != nil {
		log.Fatal("Error fetching the previous location area:", err)
	}
	
	// must update the URLs
	conf.nextURL = locationArea.Next
	conf.previousURL = locationArea.Previous
	
	//logging the previous location areas to the console
	fmt.Println("Display the previous 20 sites here...")
	pokeapi.LogLocationAreas(locationArea)
	return nil
}