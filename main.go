package main

func main() {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	c := &config{
		commands: getCommands(),
		nextURL: &baseURL,
		previousURL: nil,
	}
	startRepl(c)
}

