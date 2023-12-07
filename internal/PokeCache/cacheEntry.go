/* DocuString: This file creates a cache to store data that has been read previously for quicker output*/

package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheInfo map[string]cacheEntry
	mutex     sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// creates new cache and stores in cacheInfo map
func newCache() error {

	return nil
}

// adds new entry to cache
func add(key string, val []byte) error {

	return nil
}

// gets and entry from cache
func get(key string) ([]byte, bool, error) {
	val := make([]byte, 5)
	return val, true, nil
}

// removes all entries inside cacheInfo map
func reapLoop() error {

	return nil
}
