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

