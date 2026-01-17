package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mapCache map[string]cacheEntry
	mu       *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	var myCache = Cache{
		mapCache: make(map[string]cacheEntry),
		mu:       &sync.Mutex{},
	}
	go myCache.reapLoop(interval)
	return myCache
}

func (c *Cache) CacheSize() int {
	return len(c.mapCache)
}

func (c *Cache) Add(key string, value []byte) {
	caEntry := cacheEntry{time.Now(), value}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.mapCache[key] = caEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	retVal, ok := c.mapCache[key]
	return retVal.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, value := range c.mapCache {
		if value.createdAt.Before(now.Add(-interval)) {
			delete(c.mapCache, key)
		}
	}
}
