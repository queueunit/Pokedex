package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	apiCalls "github.com/queueunit/Pokedex/internal/api"
)

func commandCatch(state *config, opts ...string) error {

	if len(opts) == 1 {
		fmt.Print("Need a pokemon name with catch command")
		return errors.New("")
	}

	// Apparently the most difficult to catch pokemon is Blissey with a
	// exp value of 635. We will draw a random integer between 0 & 699,
	// if the number is greater than the exp, it will result in a catch

	x, _ := apiCalls.GetPokemon(opts[1], state.Cache)

	if x.BaseExperience == 0 {
		fmt.Printf("Invalid pokemon")
		return errors.New("")
	}

	catchProbability := rand.Intn(700)
	isCaught := false

	if catchProbability > x.BaseExperience {
		isCaught = true
	}

	fmt.Printf("Throwing a Pokeball at %v.\n", opts[1])
	time.Sleep(2 * time.Second)
	if isCaught {
		fmt.Printf("%v was caught!", opts[1])
		state.Pokedex[x.Name] = x
	} else {
		fmt.Printf("%v escaped!", opts[1])
	}

	return errors.New("")
}
