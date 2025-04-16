package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"

	"github.com/robgilliam/pokedex/internal/pokecache"
)

type config struct {
	cache   pokecache.Cache
	nextUrl string
	prevUrl string
	pokedex map[string]any
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func commandHelp(*config, string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}

	return nil
}

func commandCatch(conf *config, pokemonName string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	data, cached := conf.cache.Get(url)
	if !cached {
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("GET pokemon failed: %w", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Could not read pokemon response: %w", err)
		}
		conf.cache.Add(url, data)
	}

	pokemon := struct {
		BaseExperience int `json:"base_experience"`
	}{}

	if err := json.Unmarshal(data, &pokemon); err != nil {
		return fmt.Errorf("Could not process pokemon data: %w", err)
	}

	chanceToCatch := 200
	if pokemon.BaseExperience >= chanceToCatch {
		chanceToCatch = 1
	} else if pokemon.BaseExperience > 0 {
		chanceToCatch -= pokemon.BaseExperience
	}

	// fmt.Printf("DEBUG: %s has base experience %d (chance to catch: %d)\n", pokemonName, pokemon.BaseExperience, chanceToCatch)

	if rand.Intn(200) < chanceToCatch {
		fmt.Println(pokemonName + " was caught!")
		conf.pokedex[pokemonName] = true
	} else {
		fmt.Println(pokemonName + " escaped!")
	}

	fmt.Println("DEBUG: pokedex contents", conf.pokedex)

	return nil
}

func commandExplore(conf *config, locationName string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + locationName

	data, cached := conf.cache.Get(url)
	if !cached {
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("GET location failed: %w", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Could not read location response: %w", err)
		}
		conf.cache.Add(url, data)
	}

	location := struct {
		PokemonEncounters []struct {
			Pokemon struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"pokemon"`
		} `json:"pokemon_encounters"`
	}{}

	if err := json.Unmarshal(data, &location); err != nil {
		return fmt.Errorf("Could not process location data: %w", err)
	}

	for _, pokemon := range location.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}

func commandMap(conf *config, _ string) error {
	next, prev, err := doMap(conf.nextUrl, &conf.cache)

	if err == nil {
		conf.nextUrl = next
		conf.prevUrl = prev
	}

	return err
}

func commandMapb(conf *config, _ string) error {
	if conf.prevUrl == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	next, prev, err := doMap(conf.prevUrl, &conf.cache)

	if err == nil {
		conf.nextUrl = next
		conf.prevUrl = prev
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
			return "", "", fmt.Errorf("Could not read locations response: %w", err)
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
		return "", "", fmt.Errorf("Could not process locations data: %w", err)
	}

	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}

	return locations.Next, locations.Previous, nil
}

func commandExit(*config, string) error {
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
		"catch": {
			name:        "catch <pokemon>",
			description: "Attempt to catch a Pokemon",
			callback:    commandCatch,
		},
		"explore": {
			name:        "explore <area>",
			description: "Shows pokemon in the area",
			callback:    commandExplore,
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
