package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	ce       map[string]cacheEntry
	mu       *sync.RWMutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
