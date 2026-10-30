package main

import (
	"strings"
)

func cleanInput(text string) []string {

	lowerResult := strings.ToLower(text)
	result := strings.Fields(lowerResult)

	return result
}
