package pokecache

import (
	"sync"
	"time"
)

type cachEntry struct {
	createdAt time.Time
	val       []byte
}

type ResponseCache struct {
	cacheMap map[string]cachEntry
	mu       *sync.Mutex
}

func NewCache(interval time.Duration) ResponseCache {
	c := ResponseCache{
		cacheMap: make(map[string]cachEntry),
		mu:       &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (cache *ResponseCache) Add(key string, value []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.cacheMap[key] = cachEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

func (cache *ResponseCache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	value, exist := cache.cacheMap[key]
	if !exist {
		return nil, false
	}

	return value.val, true
}

func (cache *ResponseCache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	// Ranging over C(time.Time type channel) of time.Ticker struct's instance
	for range ticker.C {
		cache.reap(time.Now().UTC(), interval)
	}
}

func (cache *ResponseCache) reap(currentTime time.Time, last time.Duration) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	for key, value := range cache.cacheMap {
		if value.createdAt.Before(currentTime.Add(-last)) {
			delete(cache.cacheMap, key)
		}
	}
}
