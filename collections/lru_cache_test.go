package collections

import "testing"

func TestLRUCache_Get(t *testing.T) {
	lru := NewLRUCache(2)

	lru.Put(1, 1)
	lru.Put(2, 2)

	if lru.Get(1) != 1 {
		t.Error("Expected value 1 for key 1, but not found")
	}

	if lru.Get(3) != -1 {
		t.Error("Expected value -1 for key 3, but found")
	}
}

func TestLRUCache_Put(t *testing.T) {
	lru := NewLRUCache(2)

	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)

	if lru.Get(1) != -1 {
		t.Error("Expected value -1 for key 1, but found")
	}

	if lru.Get(2) != 2 {
		t.Error("Expected value 2 for key 2, but not found")
	}

	if lru.Get(3) != 3 {
		t.Error("Expected value 3 for key 3, but not found")
	}
}
