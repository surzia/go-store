package collections

import "errors"

type Stack struct {
	items []interface{} // 使用切片来存储栈的元素
}

// Push 将元素推入栈
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop 弹出栈顶元素
func (s *Stack) Pop() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, errors.New("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// Peek 查看栈顶元素但不弹出
func (s *Stack) Peek() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// Size 返回栈的大小
func (s *Stack) Size() int {
	return len(s.items)
}

// IsEmpty 检查栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
