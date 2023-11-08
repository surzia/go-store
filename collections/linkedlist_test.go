package collections

import (
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	ll := LinkedList{}

	ll.Append(1)
	ll.Append(2)

	if ll.head.value != 1 {
		t.Errorf("Expected head value: 1, got: %v", ll.head.value)
	}
	if ll.head.next.value != 2 {
		t.Errorf("Expected next value: 2, got: %v", ll.head.next.value)
	}
}

func TestLinkedList_Delete(t *testing.T) {
	ll := LinkedList{}

	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	err := ll.Delete(2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if ll.head.next.value != 3 {
		t.Errorf("Expected head's next value: 3, got: %v", ll.head.next.value)
	}

	err = ll.Delete(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if ll.head.value != 3 {
		t.Errorf("Expected head value: 3, got: %v", ll.head.value)
	}

	err = ll.Delete(4)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
