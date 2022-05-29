package cache

import (
	"fmt"
	"sync"
	"time"
)

type Item struct {
	Value      string
	Expiration int64
}

type Cache struct {
	items map[string]Item
	mu    sync.RWMutex
}

func (item Item) Expired() bool {
	if item.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > item.Expiration
}

func NewCache() Cache {
	return Cache{items: map[string]Item{}, mu: sync.RWMutex{}}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	item, found := c.items[key]
	if !found {
		c.mu.RUnlock()
		return "", false
	}

	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			c.mu.RUnlock()
			return "", false
		}
	}

	c.mu.RUnlock()
	return item.Value, true
}

func (c *Cache) Put(key, value string) {
	c.mu.Lock()

	_, found := c.items[key]

	if found {
		fmt.Printf("The value for key %s already exists and will be overwritten!", key)
	}

	c.items[key] = Item{
		Value:      value,
		Expiration: 0,
	}
	c.mu.Unlock()
}

func (c *Cache) Keys() []string {

	keys := make([]string, 0, len(c.items))

	for k, i := range c.items {

		if i.Expired() {
			continue
		}

		keys = append(keys, k)
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {

	currentTime := time.Now()

	if deadline.After(currentTime) {

		c.mu.Lock()

		_, found := c.items[key]

		if found {
			fmt.Printf("The value for key %s already exists and will be overwritten!", key)
		}

		c.items[key] = Item{
			Value:      value,
			Expiration: deadline.UnixNano(),
		}
		c.mu.Unlock()

	}

}

func (c *Cache) Delete(key string) {
	if c.items != nil {
		_, found := c.items[key]
		if found {
			delete(c.items, key)
		}
	}
}
