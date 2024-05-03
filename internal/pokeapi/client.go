package pokeapi

import (
	"net/http"
	"time"

	"github.com/nabin3/pokedexcli/internal/pokecache"
)

// Client blueprint, this client holds our cache and http_client
type Client struct {
	cache      pokecache.ResponseCache
	httpClient http.Client
}

// This func create and return a new instance of Client
func NewClient(timeOut, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval), // Crates a cache
		httpClient: http.Client{ // Creates a httpClient with a specified timeOut duration
			Timeout: timeOut,
		},
	}
}
