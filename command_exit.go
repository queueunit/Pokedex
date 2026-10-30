package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(state *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("")
}
