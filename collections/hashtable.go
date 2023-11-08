package collections

type HashTable struct {
	items map[string]interface{}
}

// 创建一个新的哈希表
func NewHashTable() *HashTable {
	h := &HashTable{
		items: make(map[string]interface{}),
	}
	return h
}

// 设置键值对
func (h *HashTable) Set(key string, value interface{}) {
	h.items[key] = value
}

// 获取键对应的值
func (h *HashTable) Get(key string) (interface{}, bool) {
	value, found := h.items[key]
	return value, found
}

// 删除特定键的键值对
func (h *HashTable) Delete(key string) {
	delete(h.items, key)
}

// 检查键是否存在
func (h *HashTable) Exists(key string) bool {
	_, found := h.items[key]
	return found
}
