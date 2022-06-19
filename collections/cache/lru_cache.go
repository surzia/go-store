package cache

import (
	"container/list"
	"sync"
)

// LRUCache is struct of simple least recently used cache, it has own
// size, a list that store pointer and a map that store key value pair.
// In this case, value must be integer which larger than 0. LRUCache
// returns value that corresponing to key when cache hits, otherwise
// returns -1.
type LRUCache struct {
	sync.RWMutex
	size    int
	list    *list.List
	element map[K]*list.Element
}

func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		size:    size,
		list:    new(list.List),
		element: make(map[K]*list.Element, size),
	}
}

// Get the cache value for assigned key, also move the list pointer to
// the front of the list
func (c *LRUCache) Get(key K) V {
	c.Lock()
	defer c.Unlock()

	if node, ok := c.element[key]; ok {
		// hit LRU cache, fetch key value pair
		value := node.Value.(*list.Element).Value.(Cache).Value
		// move node to the front of list
		c.list.MoveToFront(node)

		return value
	}

	return V(-1)
}

// Put the key value pair into LRU cache, Check if the LRU overflows,
// if so remove the last element of the list
func (c *LRUCache) Put(key K, value V) {
	c.Lock()
	defer c.Unlock()

	if node, ok := c.element[key]; ok {
		// there is already key in LRU cache, move node to the front of list
		c.list.MoveToFront(node)
		// update key value pair
		node.Value.(*list.Element).Value = Cache{Key: key, Value: value}
	} else {
		// there is no key in LRU cache
		if c.list.Len() == c.size {
			// LRU is full
			lastKey := c.list.Back().Value.(*list.Element).Value.(Cache).Key
			delete(c.element, lastKey)
			c.list.Remove(c.list.Back())
		}

		node := &list.Element{
			Value: Cache{
				Key:   key,
				Value: value,
			},
		}

		pointer := c.list.PushFront(node)
		c.element[key] = pointer
	}
}

// RecentlyUsed returns recently used value
func (c *LRUCache) RecentlyUsed() V {
	c.Lock()
	defer c.Unlock()

	return c.list.Front().Value.(*list.Element).Value.(Cache).Value
}

// Remove key value pair in LRU cache if key exists
func (c *LRUCache) Remove(key K) {
	c.Lock()
	defer c.Unlock()

	if node, ok := c.element[key]; ok {
		delete(c.element, key)
		c.list.Remove(node)
	}
}

// Clear the LRU cache
func (c *LRUCache) Clear() {
	c.Lock()
	defer c.Unlock()

	c.list = nil
	c.element = nil
}

// Keys print all key-value pair in list
func (c *LRUCache) Keys() []K {
	c.Lock()
	defer c.Unlock()

	var ret []K
	for k := range c.element {
		ret = append(ret, k)
	}

	return ret
}
