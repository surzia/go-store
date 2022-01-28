package queue

import "sync"

type T int

type Queue struct {
	sync.RWMutex
	array []T
}

func NewQueue() *Queue {
	queue := &Queue{}
	queue.array = []T{}

	return queue
}

// Enqueue adds t to the end of the queue
func (q *Queue) Enqueue(t T) {
	q.Lock()
	q.array = append(q.array, t)
	q.Unlock()
}

// Dequeue removes the start of from the queue
func (q *Queue) Dequeue() *T {
	q.Lock()
	ret := q.array[0]
	q.array = q.array[1:len(q.array)]
	q.Unlock()

	return &ret
}

// Front returns the start in the queue, without removing it
func (q *Queue) Front() *T {
	q.Lock()
	ret := q.array[0]
	q.Unlock()

	return &ret
}

// Size returns the size of the queue
func (q *Queue) Size() int {
	q.RLock()
	defer q.RUnlock()

	return len(q.array)
}

// IsEmpty returns true if the queue is empty
func (q *Queue) IsEmpty() bool {
	q.RLock()
	defer q.RUnlock()

	return len(q.array) == 0
}
