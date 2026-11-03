package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(state *config, opts ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("")
}
