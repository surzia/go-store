package collections

import (
	"errors"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	queue := Queue{}

	queue.Enqueue(1)
	queue.Enqueue(2)

	if queue.Size() != 2 {
		t.Errorf("Expected queue size: 2, got: %d", queue.Size())
	}
}

func TestQueue_Dequeue(t *testing.T) {
	queue := Queue{}

	queue.Enqueue(1)
	queue.Enqueue(2)

	dequeued, _ := queue.Dequeue()
	if dequeued != 1 {
		t.Errorf("Expected dequeued element: 1, got: %v", dequeued)
	}

	if queue.Size() != 1 {
		t.Errorf("Expected queue size: 1, got: %d", queue.Size())
	}
}

func TestQueue_DequeueEmptyQueue(t *testing.T) {
	queue := Queue{}

	_, err := queue.Dequeue()
	if err == nil {
		t.Error("Expected error for dequeuing from empty queue, but got nil")
	}

	expectedErr := errors.New("queue is empty")
	if err.Error() != expectedErr.Error() {
		t.Errorf("Expected error: %q, got: %q", expectedErr, err)
	}
}

func TestQueue_Front(t *testing.T) {
	queue := Queue{}

	queue.Enqueue(1)
	queue.Enqueue(2)

	front, _ := queue.Front()
	if front != 1 {
		t.Errorf("Expected front element: 1, got: %v", front)
	}

	if queue.Size() != 2 {
		t.Errorf("Expected queue size: 2, got: %d", queue.Size())
	}
}

func TestQueue_FrontEmptyQueue(t *testing.T) {
	queue := Queue{}

	_, err := queue.Front()
	if err == nil {
		t.Error("Expected error for checking front in empty queue, but got nil")
	}

	expectedErr := errors.New("queue is empty")
	if err.Error() != expectedErr.Error() {
		t.Errorf("Expected error: %q, got: %q", expectedErr, err)
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	queue := Queue{}

	if !queue.IsEmpty() {
		t.Error("Expected queue to be empty, but it's not")
	}

	queue.Enqueue(1)

	if queue.IsEmpty() {
		t.Error("Expected queue not to be empty, but it is")
	}
}
