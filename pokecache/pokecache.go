package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	interval time.Duration
	mu       sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{entries: make(map[string]cacheEntry), interval: interval}
	go reapLoop(c)
	return c
}

func reapLoop(c *Cache) {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.reap()
	}
}

func (c *Cache) reap() {
	now := time.Now()
	for key, entry := range c.entries {
		if now.Sub(entry.createdAt) > c.interval {
			c.mu.Lock()
			delete(c.entries, key)
			c.mu.Unlock()
		}
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	entry, ok := c.entries[key]
	c.mu.RUnlock()
	return entry.val, ok
}

func (c *Cache) Set(key string, val []byte) {
	c.mu.Lock()
	c.entries[key] = cacheEntry{createdAt: time.Now().UTC(), val: val}
	c.mu.Unlock()
}
