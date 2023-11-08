package collections

import "testing"

func TestRingBuffer_EnqueueDequeue(t *testing.T) {
	rb := NewRingBuffer(3)

	err := rb.Enqueue(1)
	if err != nil {
		t.Errorf("Unexpected error on Enqueue: %v", err)
	}

	item, err := rb.Dequeue()
	if err != nil || item != 1 {
		t.Errorf("Unexpected error or result on Dequeue: %v, %v", item, err)
	}
}

func TestRingBuffer_FullAndEmpty(t *testing.T) {
	rb := NewRingBuffer(2)

	if !rb.IsEmpty() {
		t.Error("Expected buffer to be empty, but it's not")
	}

	rb.Enqueue(1)
	rb.Enqueue(2)

	if !rb.IsFull() {
		t.Error("Expected buffer to be full, but it's not")
	}

	_, _ = rb.Dequeue()
	_, _ = rb.Dequeue()

	if !rb.IsEmpty() {
		t.Error("Expected buffer to be empty, but it's not")
	}
}
