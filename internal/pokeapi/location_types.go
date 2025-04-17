package pokeapi

type LocationsList struct {
	Count    int
	Next     string
	Previous string
	Results  []struct {
		Name string
		Url  string
	}
}

type Location struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string
			Url  string
		}
	} `json:"pokemon_encounters"`
}
