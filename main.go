package main

func main() {
	state := config{state: "begun",
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: "null",
	}
	startRepl(&state)
}
