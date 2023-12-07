/* DocuString: This file creates a cache to store data that has been read previously for quicker output*/

package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheInfo map[string]cacheEntry
	mux       *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// creates new cache and tracks time before deleting all
func newCache(holdInCache time.Duration) Cache {
	c := Cache{
		cacheInfo: make(map[string]cacheEntry),
		mux:       &sync.Mutex{},
	}

	go c.reapLoop(holdInCache)

	return c
}

// adds new entry to cache
func (c *Cache) add(key string, info []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cacheInfo[key] = cacheEntry{
		createdAt: time.Now(),
		val:       info,
	}
}

// gets and entry from cache
func (c *Cache) get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	val, ok := c.cacheInfo[key]
	return val.val, ok
}

// loops through cache
func (c *Cache) reapLoop(holdInCache time.Duration) {
	timeTick := time.NewTicker(holdInCache)
	for range timeTick.C {
		c.reapItem(time.Now().UTC(), holdInCache)
	}
}

// removes item that is passed
func (c *Cache) reapItem(now time.Time, holdInCache time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	for i, r := range c.cacheInfo {
		if r.createdAt.Before(now.Add(-holdInCache)) {
			delete(c.cacheInfo, i)
		}
	}
}
