package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheMap map[string]cacheEntry
	mu       sync.Mutex
}

// Create a new cache that will store data for a time the wipe.
func NewCache(interval time.Duration) *Cache {
	cMap := make(map[string]cacheEntry)
	newCache := &Cache{
		cacheMap: cMap,
	}
	go newCache.reapLoop(interval)
	return newCache
}

// Method to add to the cache
func (c *Cache) Add(key string, v []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       v,
	}
}

// Method to retrive from the cache
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	k, ok := c.cacheMap[key]
	if ok {
		return k.val, ok
	}
	return nil, false
}

// Method to loop through the data in the cache
// Wipe if data is older than interval
func (c *Cache) reapLoop(interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()

	for range t.C {
		c.mu.Lock()
		for key, entry := range c.cacheMap {
			if time.Since(entry.createdAt) >= interval {
				delete(c.cacheMap, key)
			}
		}
		c.mu.Unlock()
	}
}
