package pokecache

import (
	"sync"
	"time"
)

// This struct is blueprint for cached data which was retrieved at a certain time
type cachEntry struct {
	createdAt time.Time
	val       []byte
}

// This struct is blueprint for storing each cacheEntry struct instance
type ResponseCache struct {
	cacheMap map[string]cachEntry // Storage of cache, this map key is the URL, to which the retriever(funcs which make requests to pokepi endpoints) funcs made request
	mu       *sync.Mutex          // mutex used to prevent race condition
}

// This function creates a newcache, Before requesting to pokepi endpoint our retriever fincs will look at this cache storage, this cache storage also deletes expired cachEntry in a predefined time interval
func NewCache(interval time.Duration) ResponseCache {
	c := ResponseCache{
		cacheMap: make(map[string]cachEntry),
		mu:       &sync.Mutex{},
	}

	// CachEntry created before certain amount time should be deleted from cacheMap
	go c.reapLoop(interval)

	return c
}

// This func add new cachEntry to cacheMap
func (cache *ResponseCache) Add(key string, value []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.cacheMap[key] = cachEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

// Retrirving a cachEntry from cacheMap
func (cache *ResponseCache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	value, exist := cache.cacheMap[key]
	if !exist {
		return nil, false
	}

	return value.val, true
}

// This func implement a time loop which will triger reap func in each time interval
func (cache *ResponseCache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	// Ranging over C(time.Time type channel) of time.Ticker struct's instance
	for range ticker.C {
		cache.reap(time.Now().UTC(), interval)
	}
}

// This func delete expired cachEntry from cachMap
func (cache *ResponseCache) reap(currentTime time.Time, last time.Duration) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	for key, value := range cache.cacheMap {
		if value.createdAt.Before(currentTime.Add(-last)) {
			delete(cache.cacheMap, key)
		}
	}
}
