package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLOWORLD",
			expected: []string{"helloworld"},
		},
		{
			input:    "Hello porld",
			expected: []string{"hello", "porld"},
		},
	}

	passCount := 0
	failCount := 0

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
			// error and continue here
			failCount++
			t.Errorf(`----------------------
			Input:    %v
			Expecting:   %v
			Actual:      %v
			FAIL
			`, c.input, c.expected, actual)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			fmt.Println(word)
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				failCount++
				t.Errorf(`------------------
				Expecting:	%v
				Actual:		%v
				FAIL`, expectedWord, word)
			}
		}
		passCount++
		fmt.Printf(`----------------------
			Expecting:   %v
			Actual:      %v
			PASS
			`, c.expected, actual)
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
