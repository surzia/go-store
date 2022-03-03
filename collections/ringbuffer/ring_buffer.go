/*
modification history
--------------------
2022/3/3, 14:43, by song, create
*/

package ringbuffer

import "sync"

type T string

const DefaultCapacity = 8

type RingBuffer struct {
	sync.RWMutex
	data     []T
	capacity int
	read     int
	write    int
}

func NewRingBuffer(cap int) *RingBuffer {
	ring := &RingBuffer{}
	if cap < 1 {
		cap = DefaultCapacity
	}

	ring.data = make([]T, cap)
	ring.capacity = cap
	ring.read = 0
	ring.write = -1

	return ring
}

func (r *RingBuffer) Offer(t T) bool {
	if !r.IsFull() {
		next := r.write + 1
		r.data[next%r.capacity] = t
		r.write++
		return true
	}

	return false
}

func (r *RingBuffer) Poll() *T {
	if !r.IsEmpty() {
		next := r.data[r.read%r.capacity]
		r.read++
		return &next
	}

	return nil
}

func (r *RingBuffer) Size() int {
	r.Lock()
	defer r.Unlock()

	return r.write - r.read + 1
}

func (r *RingBuffer) IsEmpty() bool {
	r.Lock()
	defer r.Unlock()

	return r.write < r.read
}

func (r *RingBuffer) IsFull() bool {
	return r.Size() >= r.capacity
}
