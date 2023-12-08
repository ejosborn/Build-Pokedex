/* DocuString: This file creates a cache to store data that has been read previously for quicker output*/

package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// creates new cache and tracks time before deleting all
func NewCache(t time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}

	go c.reapLoop(t)

	return c
}

// adds new entry to cache
func (c *Cache) Add(key string, info []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache[key] = cacheEntry{
		val:       info,
		createdAt: time.Now().UTC(),
	}
}

// gets an entry from cache
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	cacheE, ok := c.cache[key]
	return cacheE.val, ok
}

// loops through cache tracking time in cache
func (c *Cache) reapLoop(t time.Duration) {
	ticker := time.NewTicker(t)
	for range ticker.C {
		c.reapItem(t)
	}
}

// deletes item in cache after its time is up
func (c *Cache) reapItem(t time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	timeNow := time.Now().UTC().Add(-t)
	for k, v := range c.cache {
		if v.createdAt.Before(timeNow) {
			delete(c.cache, k)
		}
	}
}
