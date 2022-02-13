package hashmap

import "sync"

type Key string
type Value string

type HashMap struct {
	sync.RWMutex
	hashmap map[int]Value
}

func NewHashMap() *HashMap {
	ret := &HashMap{hashmap: make(map[int]Value)}

	return ret
}

// Put item with value v and key k into the hashmap
func (h *HashMap) Put(k Key, v Value) {
	h.Lock()
	defer h.Unlock()

	i := hash(k)
	if h.hashmap == nil {
		h.hashmap = make(map[int]Value)
	}

	h.hashmap[i] = v
}

// Remove item with key k from hashmap
func (h *HashMap) Remove(k Key) bool {
	h.Lock()
	defer h.Unlock()

	i := hash(k)
	if _, ok := h.hashmap[i]; ok {
		delete(h.hashmap, i)
		return true
	}

	return false
}

// Get item with key k from the hashmap
func (h *HashMap) Get(k Key) *Value {
	h.RLock()
	defer h.RUnlock()

	i := hash(k)
	if v, ok := h.hashmap[i]; ok {
		return &v
	}
	return nil
}

// Size returns the number of the hashmap elements
func (h *HashMap) Size() int {
	h.RLock()
	defer h.RUnlock()

	return len(h.hashmap)
}
