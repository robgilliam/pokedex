package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (cc *CacheClient) GetLocation(locationName string) (Location, error) {
	var location Location

	url := baseUrl + "/location-area/" + locationName

	data, err := cc.Get(url)
	if err != nil {
		return location, fmt.Errorf("Could not get location data: %w", err)
	}

	if err := json.Unmarshal(data, &location); err != nil {
		return location, fmt.Errorf("Could not process location data: %w", err)
	}

	return location, nil
}
