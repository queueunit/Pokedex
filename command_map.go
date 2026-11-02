package main

import (
	"errors"
	"fmt"

	apiCalls "github.com/queueunit/Pokedex/internal/api"
)

func commandMap(state *config) error {
	x, _ := apiCalls.GetLocations(0)
	fmt.Println(x.Results[0].Name)
	return errors.New("")
}
