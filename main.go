package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/robgilliam/pokedex/internal/pokecache"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	conf := config{
		cache: pokecache.NewCache(5 * time.Second),
	}

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

			if err := command.callback(&conf); err != nil {
				log.Fatal(err)
			}
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
