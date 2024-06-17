package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	interval int64
	mu       sync.RWMutex
}

type cacheEntry struct {
	createdAt int64
	val       []byte
}

func NewCache(interval int64) *Cache {
	c := &Cache{entries: make(map[string]cacheEntry), interval: interval}
	go reapLoop(c)
	return c
}

func reapLoop(c *Cache) {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		fmt.Println("tick")
		c.reap()
	}
}

func (c *Cache) reap() {
	now := time.Now().Unix()
	for key, entry := range c.entries {
		if now-entry.createdAt > c.interval {
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
	c.entries[key] = cacheEntry{createdAt: time.Now().Unix(), val: val}
	c.mu.Unlock()
}
