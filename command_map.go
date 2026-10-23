package main

import (
	"fmt"
	"log"
	"github.com/yasharthdev/pokedexcli/internal/pokeapi"
)

func commandMap(conf *config) error {
	fmt.Println("Displaying 20 sites here...")
	// conf.next contains the baseURL right now, we update it later
	locationArea, err := pokeapi.GetLocationAreas(conf.nextURL) 
	if err != nil {
		log.Fatal("Error fetching location areas:", err)
	}
	
	// once the first page loades we update conf.next and conf.prev
	conf.nextURL = locationArea.Next
	conf.previousURL = locationArea.Previous
	
	pokeapi.LogLocationAreas(locationArea)
	return nil
}