package pokecache

import (
	"testing"
	"time"
)

func Test_Cache(t *testing.T) {
	cache := NewCache(5 * time.Second)

	cache.Add("key1", []byte("Dog"))
	cache.Add("key2", []byte("Cat"))

	if actual, got := cache.Get("key1"); got {
		if string(actual) != "Dog" {
			t.Errorf("Cache hit but wrong data: %s returned %s; expected %s", "key1", "Dog", string(actual))
		}
	} else {
		t.Errorf("Cache MISS: %s", "key1")
	}

	if actual, got := cache.Get("key2"); got {
		if string(actual) != "Cat" {
			t.Errorf("Cache hit but wrong data: %s returned %s; expected %s", "key2", "Cat", string(actual))
		}
	} else {
		t.Errorf("Cache MISS: %s", "key2")
	}

	if actual, got := cache.Get("key3"); got {
		t.Errorf("Unexpected cache hit for %s: returned %s", "key3", string(actual))
	}

	// Wait 6 seconds
	time.Sleep(6 * time.Second)

	if actual, got := cache.Get("key1"); got {
		t.Errorf("Unexpected cache hit for %s: returned %s", "key1", string(actual))
	}
	if actual, got := cache.Get("key2"); got {
		t.Errorf("Unexpected cache hit for %s: returned %s", "key2", string(actual))
	}
	if actual, got := cache.Get("key3"); got {
		t.Errorf("Unexpected cache hit for %s: returned %s", "key3", string(actual))
	}

}
