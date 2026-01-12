package pokecache

import (
	"sync"
	"time"
	"fmt"
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


func (c *Cache) Add(key string, value []byte) {
    c.mux.Lock()
    defer c.mux.Unlock()
    c.pokeCache[key] = cacheEntry{
        createdAt: time.Now().UTC(),
        val:       value,
    }
	fmt.Println("Added to cache!")
	
}

func (c *Cache) Get(url string) ([]byte, bool) {
	dat, ok := c.pokeCache[url]
	if ok {
		return dat.val, true}
	return dat.val, false
}

