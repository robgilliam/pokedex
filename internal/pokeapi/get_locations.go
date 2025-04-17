package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (cc *CacheClient) GetLocations(url string) (LocationsList, error) {
	var locations LocationsList

	if url == "" {
		url = baseUrl + "/location-area"
	}

	data, err := cc.Get(url)
	if err != nil {
		return locations, fmt.Errorf("Could not get locations data: %w", err)
	}

	if err := json.Unmarshal(data, &locations); err != nil {
		return locations, fmt.Errorf("Could not process locations data: %w", err)
	}

	return locations, nil
}
