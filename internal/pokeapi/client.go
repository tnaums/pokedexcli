package pokeapi

import (
	"net/http"
	"time"

	"pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	httpClient http.Client
	dataCache pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		dataCache: pokecache.NewCache(5 * time.Second),
		
	}
}
