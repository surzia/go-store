package collections

import "container/list"

type LRUCache struct {
	capacity int
	items    map[int]*list.Element
	queue    *list.List
}

type entry struct {
	key   int
	value int
}

// 创建一个新的LRU缓存
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[int]*list.Element),
		queue:    list.New(),
	}
}

// 获取LRU缓存中的值
func (lru *LRUCache) Get(key int) int {
	if elem, found := lru.items[key]; found {
		lru.queue.MoveToFront(elem)
		return elem.Value.(*entry).value
	}
	return -1
}

// 设置LRU缓存中的值
func (lru *LRUCache) Put(key, value int) {
	if elem, found := lru.items[key]; found {
		lru.queue.MoveToFront(elem)
		elem.Value.(*entry).value = value
	} else {
		if len(lru.items) >= lru.capacity {
			lru.evict()
		}
		newElem := lru.queue.PushFront(&entry{key, value})
		lru.items[key] = newElem
	}
}

// 从LRU缓存中删除最久未使用的项
func (lru *LRUCache) evict() {
	if lru.queue.Len() == 0 {
		return
	}
	removed := lru.queue.Back()
	lru.queue.Remove(removed)
	delete(lru.items, removed.Value.(*entry).key)
}
