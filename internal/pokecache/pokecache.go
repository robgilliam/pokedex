package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	c    map[string]cacheEntry
	lock sync.RWMutex
}

func NewCache(interval time.Duration) Cache {
	ch := Cache{
		c: make(map[string]cacheEntry),
	}

	ch.reapLoop(interval)

	return ch
}

func (ch *Cache) Add(key string, data []byte) {
	ch.lock.Lock()
	defer ch.lock.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val:       data,
	}
	ch.c[key] = entry
}

func (ch *Cache) Get(key string) ([]byte, bool) {
	ch.lock.RLock()
	defer ch.lock.RUnlock()

	entry, hasEntry := ch.c[key]

	data := []byte{}

	if hasEntry {
		data = entry.val
	}

	return data, hasEntry
}

func (ch *Cache) reapLoop(interval time.Duration) {
	t := time.NewTicker(interval)

	go func() {
		for {
			now := <-t.C

			ch.lock.Lock()

			for key, entry := range ch.c {
				if now.Sub(entry.createdAt) > interval {
					delete(ch.c, key)
				}
			}

			ch.lock.Unlock()
		}
	}()
}
