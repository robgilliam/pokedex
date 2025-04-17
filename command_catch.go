package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(conf *config, pokemonNames ...string) error {
	for _, pokemonName := range pokemonNames {

		fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

		pokemon, err := conf.client.GetPokemon(pokemonName)
		if err != nil {
			return fmt.Errorf("Could not get Pokemon information for %s", pokemonName)
		}

		chanceToCatch := 200
		if pokemon.BaseExperience >= chanceToCatch {
			chanceToCatch = 1
		} else if pokemon.BaseExperience > 0 {
			chanceToCatch -= pokemon.BaseExperience
		}

		if rand.Intn(200) < chanceToCatch {
			fmt.Println(pokemonName + " was caught!")
			conf.pokedex[pokemonName] = pokemon
		} else {
			fmt.Println(pokemonName + " escaped!")
		}
	}

	return nil
}
