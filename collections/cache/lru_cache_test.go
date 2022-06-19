package cache

import (
	"container/list"
	"testing"
)

var (
	key1 K = 1
	key2 K = 2
	key3 K = 3
	key4 K = 4

	value1 V = 5
	value2 V = 6
	value3 V = 7
	value4 V = 8
)

func InitLRUCache() *LRUCache {
	cache := NewLRUCache(3)
	node1 := &list.Element{
		Value: Cache{
			Key:   key1,
			Value: value1,
		},
	}
	pointer1 := cache.list.PushFront(node1)
	cache.element[key1] = pointer1

	node2 := &list.Element{
		Value: Cache{
			Key:   key2,
			Value: value2,
		},
	}
	pointer2 := cache.list.PushFront(node2)
	cache.element[key2] = pointer2
	return cache
}

func TestGet(t *testing.T) {
	cache := InitLRUCache()
	v1 := cache.Get(key1)
	if v1 != value1 {
		t.Errorf("wrong count, expected %d and got %d", value1, v1)
	}
	v2 := cache.Get(key2)
	if v2 != value2 {
		t.Errorf("wrong count, expected %d and got %d", value2, v2)
	}
}

func TestPut(t *testing.T) {
	cache := InitLRUCache()
	cache.Put(key3, value3)
	v3 := cache.Get(key3)
	if v3 != value3 {
		t.Errorf("wrong count, expected %d and got %d", value3, v3)
	}

	cache.Put(key4, value4)

	v1 := cache.Get(key1)
	if v1 != -1 {
		t.Errorf("wrong count, expected %d and got %d", -1, v1)
	}

	lru := cache.RecentlyUsed()
	if lru != value4 {
		t.Errorf("wrong count, expected %d and got %d", value4, lru)
	}

	v4 := cache.Get(key4)
	if v4 != value4 {
		t.Errorf("wrong count, expected %d and got %d", value4, v4)
	}

	cache.Get(key2)
	lru = cache.RecentlyUsed()
	if lru != value2 {
		t.Errorf("wrong count, expected %d and got %d", value2, lru)
	}

	cache.Put(key3, value1)
	lru = cache.RecentlyUsed()
	if lru != value1 {
		t.Errorf("wrong count, expected %d and got %d", value1, lru)
	}

	v3 = cache.Get(key3)
	if v3 != value1 {
		t.Errorf("wrong count, expected %d and got %d", value1, v3)
	}
}
