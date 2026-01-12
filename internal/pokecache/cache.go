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
	x := Cache{
		pokeCache: make(map[string]cacheEntry),
		mux: sync.Mutex{},
	}
	go x.reapLoop(interval)
	return x
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
	c.mux.Lock()
	dat, ok := c.pokeCache[url]
	c.mux.Unlock()
	if ok {
		fmt.Println("Found in cache!")
		return dat.val, true}
	fmt.Println("Nothing in cache!")
	return dat.val, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.pokeCache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.pokeCache, k)
		}
	}
}
