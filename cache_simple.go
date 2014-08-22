package cache

import "sync"

type CacheSimple struct {
	mu   sync.Mutex
	data map[string]string
}

func NewCacheSimple() *CacheSimple {
	return &CacheSimple{data: make(map[string]string, 64)}
}

func (c *CacheSimple) Write(key, value string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
	return nil
}

func (c *CacheSimple) Read(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v, ok := c.data[key]; ok {
		return v, true
	}
	return "", false
}
