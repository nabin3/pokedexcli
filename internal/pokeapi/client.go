package pokeapi

import (
	"net/http"
	"time"

	"github.com/nabin3/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	cache      pokecache.ResponseCache
	httpClient http.Client
}

// NewClient -
func NewClient(timeOut, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeOut,
		},
	}
}
