package main

import (
	"fmt"
)

func commandMap(conf *config, _ ...string) error {
	return doMap(conf, conf.nextLocationsUrl)
}

func commandMapb(conf *config, _ ...string) error {
	if conf.prevLocationsUrl == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	return doMap(conf, conf.prevLocationsUrl)
}

func doMap(conf *config, url string) error {
	locations, err := conf.client.GetLocations(url)
	if err != nil {
		return fmt.Errorf("Could not get locations information")
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	conf.nextLocationsUrl = locations.Next
	conf.prevLocationsUrl = locations.Previous

	return nil
}
