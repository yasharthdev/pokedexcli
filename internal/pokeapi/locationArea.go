package pokeapi

import "fmt"

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func LogLocationAreas(locationArea LocationArea) {
	// note that locationArea.Results is a slice of structs
	// each containing a Name and a URL
	for _, result := range locationArea.Results {
		fmt.Printf("Name: %s, URL: %s\n", result.Name, result.URL)		
	}
}