package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		input := cleanInput(scanner.Text())
		if len(input) > 0 {
			command, exists := getCommands()[input[0]]
			if !exists {
				fmt.Println("Unknown command")
				continue
			}

			command.callback()

		}
	}
}

func cleanInput(text string) []string {
	var r []string

	for s := range strings.SplitSeq(text, " ") {
		s = strings.TrimSpace(s)

		if s != "" {
			r = append(r, strings.ToLower(s))
		}
	}

	return r
}
