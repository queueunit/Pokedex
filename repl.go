package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	apiCalls "github.com/queueunit/Pokedex/internal/api"
	pokecache "github.com/queueunit/Pokedex/internal/cache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	state    string
	Next     string
	Previous string
	Cache    *pokecache.Cache
	Pokedex  map[string]apiCalls.Pokemon
}

func cleanInput(text string) []string {

	lowerResult := strings.ToLower(text)
	result := strings.Fields(lowerResult)

	return result
}

func startRepl(state *config) {

	mapping := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays help options",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Gets the 20 available locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Gets the previous 20 available locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Takes a region name and returns the pokemon in that area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a pokemon and add to the users collection",
			callback:    commandCatch,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		line := scanner.Text()
		cleanInput := cleanInput(line)
		if len(cleanInput) == 0 {
			continue
		}

		x, exists := mapping[cleanInput[0]]
		if exists {
			if len(cleanInput) > 1 {
				//				fmt.Println("Variadic function call")
				err := x.callback(state, cleanInput[0:]...)
				fmt.Println(err)
			} else {
				err := x.callback(state, cleanInput[0])
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

		//		fmt.Println("Your command was:", cleanInput[0])
		err := scanner.Err()
		if err != nil {
			fmt.Println("Something went way wrong...")
		}
		continue
	}
}
