package collections

import "errors"

type Queue struct {
	items []interface{}
}

// Enqueue 将元素加入队列尾部
func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

// Dequeue 弹出队列头部元素
func (q *Queue) Dequeue() (interface{}, error) {
	if len(q.items) == 0 {
		return nil, errors.New("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

// Front 返回队列头部元素但不弹出
func (q *Queue) Front() (interface{}, error) {
	if len(q.items) == 0 {
		return nil, errors.New("queue is empty")
	}
	return q.items[0], nil
}

// Size 返回队列大小
func (q *Queue) Size() int {
	return len(q.items)
}

// IsEmpty 检查队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
