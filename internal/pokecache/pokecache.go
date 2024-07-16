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
	store map[string]cacheEntry
	mu    *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		store: map[string]cacheEntry{},
		mu:    &sync.Mutex{},
	}

	go cache.readLoop(interval)

	return cache
}

func (cache Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.store[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (cache Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	entry, ok := cache.store[key]
	return entry.val, ok
}

func (cache Cache) readLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	func() {
		for range ticker.C {
			now := time.Now().UTC()
			cache.mu.Lock()
			for k, v := range cache.store {
				if now.After(v.createdAt.Add(interval)) {
					delete(cache.store, k)
				}
			}
			cache.mu.Unlock()
		}
	}()
}
