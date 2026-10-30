package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {

	lowerResult := strings.ToLower(text)
	result := strings.Fields(lowerResult)

	return result
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		line := scanner.Text()
		cleanInput := cleanInput(line)
		fmt.Println("Your command was:", cleanInput[0])
		continue
	}
	//	for scanner.Scan() {
	//		fmt.Println("Pokedex >")
	//		line := scanner.Text()
	//		cleanInput := cleanInput(line)
	//		fmt.Println("Your command was: ", cleanInput)
	//		continue
	//	}
	//
	//	if err := scanner.Err(); err != nil {
	//		errors.New("Error reading input")
	//	}
}
