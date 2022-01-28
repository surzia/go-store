package stack

import "testing"

var (
	t1 T = 1
	t2 T = 2
	t3 T = 3
)

func TestStack_Push(t *testing.T) {
	stack := NewStack()
	stack.Push(t1)
	stack.Push(t2)
	stack.Push(t3)

	first := stack.array[0]
	last := stack.array[len(stack.array)-1]

	if first != t1 || last != t3 {
		t.Errorf("wrong order, expected first 1 and last 3 but got %d and %d", t1, t3)
	}
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack()
	stack.Push(t1)
	stack.Push(t2)
	stack.Push(t3)

	_, _ = stack.Pop()
	if size := stack.Size(); size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}
	_, _ = stack.Pop()
	_, _ = stack.Pop()
	if size := stack.Size(); size != 0 {
		t.Errorf("wrong count, expected 0 and got %d", size)
	}

	_, err := stack.Pop()
	if err == nil {
		t.Errorf("stack must not be empty")
	}
}

func TestStack_Size(t *testing.T) {
	stack := NewStack()
	stack.Push(t1)
	stack.Push(t2)
	stack.Push(t3)

	if size := stack.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
}

func TestStack_IsEmpty(t *testing.T) {
	stack := NewStack()
	empty := stack.IsEmpty()
	if !empty {
		t.Errorf("wrong status, expected true and got %t", empty)
	}
	stack.Push(t1)
	empty = stack.IsEmpty()
	if empty {
		t.Errorf("wrong status, expected false and got %t", empty)
	}
}
