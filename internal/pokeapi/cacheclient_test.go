package pokeapi

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type CatFact struct {
	Fact   string
	Length int
}

func TestCacheClient(t *testing.T) {
	cc := NewCacheClient(5*time.Second, 2*time.Second)

	fact1, err := getCatFact(t, cc)
	if err != nil {
		t.Error(err)
	}

	fact, err := getCatFact(t, cc)
	if err != nil {
		t.Error(err)
	}
	assert.EqualValues(t, fact1, fact, "Fact was not cached")

	time.Sleep(4 * time.Second)

	fact, err = getCatFact(t, cc)
	if err != nil {
		t.Error(err)
	}
	assert.NotEqualValues(t, fact1, fact, "Cached fact was not expunged")
}

func getCatFact(t *testing.T, cc *CacheClient) (CatFact, error) {
	data, err := cc.Get("https://catfact.ninja/fact")

	if err != nil {
		return CatFact{}, fmt.Errorf("Error getting data: %w", err)
	}

	var catFact CatFact
	if err := json.Unmarshal(data, &catFact); err != nil {
		return CatFact{}, fmt.Errorf("Error unmarshalling data: %w", err)
	}

	return catFact, nil
}
