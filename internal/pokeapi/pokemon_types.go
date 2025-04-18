package pokeapi

import "fmt"

type Pokemon struct {
	Name           string
	Height         int
	Weight         int
	BaseExperience int `json:"base_experience"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string
		}
	}
	Types []struct {
		Type struct {
			Name string
		}
	}
}

func (p Pokemon) String() string {
	const pokemonFmt = `Name: %s
Height: %d
Weight: %d
Stats:%s
Types:%s`

	// Stats
	var stats string
	for _, s := range p.Stats {
		stats += fmt.Sprintf("\n- %s: %d", s.Stat.Name, s.BaseStat)
	}

	// Types
	var types string
	for _, t := range p.Types {
		types += fmt.Sprintf("\n- %s", t.Type.Name)
	}

	return fmt.Sprintf(pokemonFmt, p.Name, p.Height, p.Weight, stats, types)
}
