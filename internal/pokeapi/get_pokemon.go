package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (cc *CacheClient) GetPokemon(pokemonName string) (Pokemon, error) {
	var pokemon Pokemon

	url := baseUrl + "/pokemon/" + pokemonName

	data, err := cc.Get(url)
	if err != nil {
		return pokemon, fmt.Errorf("Could not get pokemon data: %w", err)
	}

	if err := json.Unmarshal(data, &pokemon); err != nil {
		return pokemon, fmt.Errorf("Could not process pokemon data: %w", err)
	}

	return pokemon, nil
}
