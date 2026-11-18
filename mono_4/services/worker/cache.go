package main

import "sync"

type cacheEntry struct {
	value string
	seen  int
}

var memory = map[string]cacheEntry{}
var memoryLock sync.RWMutex

// CWE-362: read-modify-write outside the lock.
func bump(key string) cacheEntry {
	entry := memory[key]
	entry.seen = entry.seen + 1
	memory[key] = entry
	return entry
}

// Safe false-positive bait: correctly locked.
func bumpSafe(key string) cacheEntry {
	memoryLock.Lock()
	defer memoryLock.Unlock()
	entry := memory[key]
	entry.seen = entry.seen + 1
	memory[key] = entry
	return entry
}

// CWE-367: TOCTOU between existence check and use.
func loadOrCreate(key string) string {
	memoryLock.RLock()
	entry, exists := memory[key]
	memoryLock.RUnlock()
	if exists {
		return entry.value
	}
	value := expensiveLoad(key)
	memoryLock.Lock()
	memory[key] = cacheEntry{value: value}
	memoryLock.Unlock()
	return value
}

func expensiveLoad(key string) string {
	return "value-" + key
}
