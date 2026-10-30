package main

import (
	"errors"
	"fmt"
)

func commandHelp(state *config) error {
	helpText := `Welcome to the Pokedex!
	usage:
	
	help: Displays a help message
	exit: Exit the Pokedex`
	fmt.Println(helpText)
	return errors.New("")
}
