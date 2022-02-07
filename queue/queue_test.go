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
	_, ret := queue.Dequeue()
	if *ret != t1 {
		t.Errorf("wrong result, expected %d and got %d", *ret, t1)
	}

	_, _ = queue.Dequeue()
	_, ret = queue.Dequeue()
	if *ret != t3 {
		t.Errorf("wrong result, expected %d and got %d", *ret, t3)
	}

	err, _ := queue.Dequeue()
	if !queue.IsEmpty() {
		t.Errorf("IsEmpty should return true")
	}
	if err == nil {
		t.Errorf("cannot operate dequeue when queue is empty")
	}
}
