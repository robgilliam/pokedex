package main

import (
	"fmt"
)

func commandInspect(conf *config, pokemonName ...string) error {

	if len(pokemonName) != 1 {
		fmt.Println("Usage: inspect <pokemon>")
		return nil
	}

	if pokemon, caught := conf.pokedex[pokemonName[0]]; caught {
		fmt.Println(pokemon)
	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}
