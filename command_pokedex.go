package main

import (
	"errors"
	"fmt"
)

func commandPokedex(state *config, opts ...string) error {
	fmt.Printf("Your Pokedex:\n")
	for _, pok := range state.Pokedex {
		fmt.Printf(" - %v\n", pok.Name)
	}
	return errors.New("")

}
