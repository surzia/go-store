package collections

import "errors"

type RingBuffer struct {
	data     []interface{}
	size     int
	capacity int
	head     int
	tail     int
}

// 创建新的环形缓冲区
func NewRingBuffer(capacity int) *RingBuffer {
	return &RingBuffer{
		data:     make([]interface{}, capacity),
		capacity: capacity,
		head:     0,
		tail:     0,
		size:     0,
	}
}

// 向环形缓冲区中添加元素
func (r *RingBuffer) Enqueue(item interface{}) error {
	if r.size == r.capacity {
		return errors.New("buffer is full")
	}
	r.data[r.tail] = item
	r.tail = (r.tail + 1) % r.capacity
	r.size++
	return nil
}

// 从环形缓冲区中移除元素
func (r *RingBuffer) Dequeue() (interface{}, error) {
	if r.size == 0 {
		return nil, errors.New("buffer is empty")
	}
	item := r.data[r.head]
	r.head = (r.head + 1) % r.capacity
	r.size--
	return item, nil
}

// 返回环形缓冲区的大小
func (r *RingBuffer) Size() int {
	return r.size
}

// 检查环形缓冲区是否为空
func (r *RingBuffer) IsEmpty() bool {
	return r.size == 0
}

// 检查环形缓冲区是否已满
func (r *RingBuffer) IsFull() bool {
	return r.size == r.capacity
}
