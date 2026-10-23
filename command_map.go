package main

import (
	"fmt"
	"log"
	"github.com/yasharthdev/pokedexcli/internal/pokeapi"
)

func commandMap(conf *config) error {
	fmt.Println("Displaying 20 sites here...")
	url := "https://pokeapi.co/api/v2/location-area/"
	
	locationArea, err := pokeapi.GetLocationAreas(url) 
	if err != nil {
		log.Fatal("Error fetching location areas:", err)
	}
	
	pokeapi.LogLocationAreas(locationArea)
	return nil
}