package main

import (
	"fmt"
	"testing"
)

func Test_CleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "one",
			expected: []string{"one"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "Pikachu BULBASAUR meowth",
			expected: []string{"pikachu", "bulbasaur", "meowth"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		fmt.Println(actual)

		for i, expectedWord := range c.expected {
			var actualWord string
			if i < len(actual) {
				actualWord = actual[i]
			}
			if actualWord != expectedWord {
				t.Errorf("Expected: '%s' - Actual: '%s'", expectedWord, actualWord)
			}
		}
	}
}
