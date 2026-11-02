package main

import (
	"errors"
	"fmt"

	apiCalls "github.com/queueunit/Pokedex/internal/api"
)

func commandMapb(state *config) error {
	if state.Previous == "" {
		fmt.Print("you're on the first page")
		return errors.New("")
	}
	x, _ := apiCalls.GetLocations(state.Previous)

	// Need all this for pretty printing, probably a better solution...
	for _, loc := range x.Results[:len(x.Results)-1] {
		fmt.Println(loc.Name)
	}
	fmt.Print(x.Results[len(x.Results)-1].Name)

	state.Next = x.Next
	state.Previous = x.Previous

	return errors.New("")
}
