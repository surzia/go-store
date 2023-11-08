package collections

type Set struct {
	items map[interface{}]struct{}
}

// 创建一个新的集合
func NewSet() *Set {
	s := &Set{
		items: make(map[interface{}]struct{}),
	}
	return s
}

// 添加元素到集合
func (s *Set) Add(item interface{}) {
	s.items[item] = struct{}{}
}

// 从集合中移除元素
func (s *Set) Remove(item interface{}) {
	delete(s.items, item)
}

// 检查元素是否存在于集合中
func (s *Set) Contains(item interface{}) bool {
	_, found := s.items[item]
	return found
}

// 返回集合大小
func (s *Set) Size() int {
	return len(s.items)
}
