package cache

import (
	"strconv"
	"sync"
)

type cache struct {
	mu        sync.RWMutex
	cachedKey map[string]bool
}

func NewCache() *cache {
	return &cache{
		cachedKey: make(map[string]bool),
	}
}

func (c *cache) add(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cachedKey[key] = true
	return nil
}

func (c *cache) get(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.cachedKey[key]; ok {
		return true
	}
	return false
}

func (c *cache) AddIntegerKey(key int) error {
	return c.add(strconv.Itoa(key))
}

func (c *cache) IsIntegerKeyExist(key int) bool {
	return c.get(strconv.Itoa(key))
}
