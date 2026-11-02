package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	state    string
	Next     string
	Previous string
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
			err := x.callback(state)
			fmt.Println(err)
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
