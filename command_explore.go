package main

import (
	"fmt"
)

func commandExplore(conf *config, locationName ...string) error {
	if len(locationName) != 1 {
		return fmt.Errorf("One (and only one) location name must be provided")
	}

	location, err := conf.client.GetLocation(locationName[0])
	if err != nil {
		return fmt.Errorf("Could not get location information")
	}

	for _, pokemon := range location.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
