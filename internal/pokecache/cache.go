package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	c := Cache{
		ce:       make(map[string]cacheEntry),
		mu:       &sync.RWMutex{},
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.ce[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	cache, ok := c.ce[key]
	if !ok {
		return []byte{}, false
	}
	return cache.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for url, cache := range c.ce {
			if time.Now().After(cache.createdAt.Add(c.interval)) {
				delete(c.ce, url)
			}
		}
		c.mu.Unlock()
	}
}
