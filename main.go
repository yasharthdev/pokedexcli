package main

func main() {
	c := &config{
		commands: getCommands(),
	}
	startRepl(c)
}

