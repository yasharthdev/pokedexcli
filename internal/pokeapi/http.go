package pokeapi

import (
	"fmt"
	"encoding/json"
	"net/http"
)

// must capitalize the function if we want to import it elsewhere
func GetLocationAreas(url string) (LocationArea, error) {
	// here we're going to make the get request for location areas
	res, err := http.Get(url)
	if err != nil {
		// return an empty LocationArea struct
		return LocationArea{}, fmt.Errorf("Error getting a response: %v", err)
	}
	defer res.Body.Close()
	
	var locationArea LocationArea
	if err := json.NewDecoder(res.Body).Decode(&locationArea); err != nil {
		return LocationArea{}, fmt.Errorf("can't decode JSON: %v", err)
	}
	
	return locationArea, nil
}
