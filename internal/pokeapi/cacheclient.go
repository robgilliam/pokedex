package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type CacheClient struct {
	cache      map[string]cacheEntry
	lock       sync.RWMutex
	httpClient http.Client
}

func NewCacheClient(timeout time.Duration, interval time.Duration) *CacheClient {
	cc := CacheClient{
		cache: make(map[string]cacheEntry),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}

	cc.reapLoop(interval)

	return &cc
}

func (cc *CacheClient) Get(url string) ([]byte, error) {
	if data, cached := cc.get(url); cached {
		return data, nil
	}

	res, err := cc.httpClient.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("Could not GET from %s error: %w", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return []byte{}, fmt.Errorf("Get returned non-OK status: %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Could not read data from response: %w", err)
	}

	cc.add(url, data)

	return data, nil
}

func (cc *CacheClient) add(key string, data []byte) {
	cc.lock.Lock()
	defer cc.lock.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val:       data,
	}
	cc.cache[key] = entry
}

func (cc *CacheClient) get(key string) ([]byte, bool) {
	cc.lock.RLock()
	defer cc.lock.RUnlock()

	entry, cached := cc.cache[key]

	data := []byte{}

	if cached {
		data = entry.val
	}

	return data, cached
}

func (cc *CacheClient) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for {
			now := <-ticker.C

			cc.lock.Lock()

			for key, entry := range cc.cache {
				if now.Sub(entry.createdAt) > interval {
					delete(cc.cache, key)
				}
			}

			cc.lock.Unlock()
		}
	}()
}
