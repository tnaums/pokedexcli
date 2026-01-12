package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	pokeCache map[string]cacheEntry
	mux sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	return Cache{
		pokeCache: make(map[string]cacheEntry),
		mux: sync.Mutex{},
	}
}

func (c *Cache) Add(url string, val []byte)  {
	newEntry := cacheEntry{time.Now(), val}

	c.pokeCache[url] = newEntry
}

