package queue

import "testing"

var (
	t1 T = 1
	t2 T = 2
	t3 T = 3
)

func TestQueue_Enqueue(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(t1)
	queue.Enqueue(t2)
	queue.Enqueue(t3)

	if size := queue.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
}

func TestQueue_Dequeue(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(t1)
	queue.Enqueue(t2)
	queue.Enqueue(t3)
	queue.Dequeue()
	if size := queue.Size(); size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}

	queue.Dequeue()
	queue.Dequeue()
	if size := queue.Size(); size != 0 {
		t.Errorf("wrong count, expected 0 and got %d", size)
	}

	if !queue.IsEmpty() {
		t.Errorf("IsEmpty should return true")
	}
}
