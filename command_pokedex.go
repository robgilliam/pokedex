package main

import (
	"fmt"
)

func commandPokedex(conf *config, pokemonName ...string) error {
	for pokemonName := range conf.pokedex {
		fmt.Println("- " + pokemonName)
	}

	return nil
}
