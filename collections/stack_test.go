package collections

import (
	"errors"
	"testing"
)

func TestStack_Push(t *testing.T) {
	stack := Stack{}

	stack.Push(1)
	stack.Push(2)

	if stack.Size() != 2 {
		t.Errorf("Expected stack size: 2, got: %d", stack.Size())
	}
}

func TestStack_Pop(t *testing.T) {
	stack := Stack{}

	stack.Push(1)
	stack.Push(2)

	popped, _ := stack.Pop()
	if popped != 2 {
		t.Errorf("Expected popped element: 2, got: %v", popped)
	}

	if stack.Size() != 1 {
		t.Errorf("Expected stack size: 1, got: %d", stack.Size())
	}
}

func TestStack_PopEmptyStack(t *testing.T) {
	stack := Stack{}

	_, err := stack.Pop()
	if err == nil {
		t.Error("Expected error for popping from empty stack, but got nil")
	}

	expectedErr := errors.New("stack is empty")
	if err.Error() != expectedErr.Error() {
		t.Errorf("Expected error: %q, got: %q", expectedErr, err)
	}
}

func TestStack_Peek(t *testing.T) {
	stack := Stack{}

	stack.Push(1)
	stack.Push(2)

	topElement, _ := stack.Peek()
	if topElement != 2 {
		t.Errorf("Expected top element: 2, got: %v", topElement)
	}

	if stack.Size() != 2 {
		t.Errorf("Expected stack size: 2, got: %d", stack.Size())
	}
}

func TestStack_PeekEmptyStack(t *testing.T) {
	stack := Stack{}

	_, err := stack.Peek()
	if err == nil {
		t.Error("Expected error for peeking in empty stack, but got nil")
	}

	expectedErr := errors.New("stack is empty")
	if err.Error() != expectedErr.Error() {
		t.Errorf("Expected error: %q, got: %q", expectedErr, err)
	}
}

func TestStack_IsEmpty(t *testing.T) {
	stack := Stack{}

	if !stack.IsEmpty() {
		t.Error("Expected stack to be empty, but it's not")
	}

	stack.Push(1)

	if stack.IsEmpty() {
		t.Error("Expected stack not to be empty, but it is")
	}
}
