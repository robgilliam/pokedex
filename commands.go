package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/robgilliam/pokedex/internal/pokecache"
)

type config struct {
	cache   pokecache.Cache
	NextUrl string
	PrevUrl string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func commandHelp(*config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}

	return nil
}

func commandMap(conf *config) error {
	next, prev, err := doMap(conf.NextUrl, &conf.cache)

	if err == nil {
		conf.NextUrl = next
		conf.PrevUrl = prev
	}

	return err
}

func commandMapb(conf *config) error {
	if conf.PrevUrl == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	next, prev, err := doMap(conf.PrevUrl, &conf.cache)

	if err == nil {
		conf.NextUrl = next
		conf.PrevUrl = prev
	}

	return err
}

func doMap(url string, cache *pokecache.Cache) (string, string, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	data, cached := cache.Get(url)
	if !cached {
		res, err := http.Get(url)
		if err != nil {
			return "", "", fmt.Errorf("GET locations failed: %w", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return "", "", fmt.Errorf("Could not read response body : %w", err)
		}
		cache.Add(url, data)
	}

	locations := struct {
		Count    int    `json:"count"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		}
	}{}

	if err := json.Unmarshal(data, &locations); err != nil {
		return "", "", fmt.Errorf("Could not read response body: %w", err)
	}

	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}

	return locations.Next, locations.Previous, nil
}

func commandExit(*config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "List next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous 20 locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
