package main

import (
	"errors"
	"fmt"
	"strings"

	apiCalls "github.com/queueunit/Pokedex/internal/api"
)

func commandExplore(state *config, opts ...string) error {
	if len(opts) == 1 {
		fmt.Print("Need area with explore command")
		return errors.New("")
	}
	x, _ := apiCalls.GetDetailedLocation(opts[1], state.Cache)

	arr := make([]string, 0)

	for _, pokemon := range x.PokemonEncounters {
		arr = append(arr, " - "+pokemon.Pokemon.Name)
	}

	if len(arr) == 0 {
		fmt.Print("No pokemon found...")
		return errors.New("")
	}

	results := strings.Join(arr, "\n")
	fmt.Print(results)

	return errors.New("")
}
