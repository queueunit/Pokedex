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
	cacheTable map[string]cacheEntry
	mu         *sync.RWMutex
	interval   time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cacheTable: map[string]cacheEntry{},
		mu:         &sync.RWMutex{},
		interval:   interval,
	}

	go cache.ReapLoop()
	return cache
}

func (cache *Cache) Add(k string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.cacheTable[k] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (cache *Cache) Get(k string) ([]byte, bool) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	x, exists := cache.cacheTable[k]
	if exists {
		return x.val, true
	} else {
		return []byte{}, false
	}
}

func (cache *Cache) ReapLoop() {
	ticker := time.NewTicker(cache.interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		cache.mu.Lock()

		for key, entry := range cache.cacheTable {
			if time.Since(entry.createdAt) > cache.interval {
				delete(cache.cacheTable, key)
			}
		}
		cache.mu.Unlock()
	}
}
