package main

import (
	"time"

	pokecache "github.com/queueunit/Pokedex/internal/cache"
)

func main() {
	state := config{state: "begun",
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
		Cache:    pokecache.NewCache(time.Duration(60) * time.Second),
	}
	startRepl(&state)
}
