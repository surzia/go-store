package linkedlist

import (
	"fmt"
	"sync"
)

// LinkedList the linked list struct
type LinkedList struct {
	sync.RWMutex
	head *Node
	size int
}

func NewLinkedList() *LinkedList {
	list := &LinkedList{}

	return list
}

// Add adds t to the end of the linked list
func (l *LinkedList) Add(t T) {
	l.Lock()
	defer l.Unlock()

	node := NewNode(t)
	if l.head == nil {
		l.head = node
	} else {
		last := l.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = node
	}
	l.size++
}

// Insert adds t at position pos
func (l *LinkedList) Insert(t T, pos int) error {
	l.Lock()
	defer l.Unlock()

	// validate position
	if pos < 0 || pos > l.size {
		return fmt.Errorf("insert position must be larger than 0 and smaller than linked list size")
	}

	node := NewNode(t)
	if pos == 0 {
		node.next = l.head
		l.head = node
		return nil
	}

	head := l.head
	idx := 0
	for idx < pos-2 {
		idx++
		head = head.next
	}

	node.next = head.next
	head.next = node
	l.size++
	return nil
}

// RemoveAt removes a node at position pos
func (l *LinkedList) RemoveAt(pos int) (*T, error) {
	l.Lock()
	defer l.Unlock()

	// validate position
	if pos < 0 || pos > l.size {
		return nil, fmt.Errorf("insert position must be larger than 0 and smaller than linked list size")
	}

	head := l.head
	idx := 0
	for idx < pos-1 {
		idx++
		head = head.next
	}
	next := head.next
	head.next = next.next
	l.size--
	return &next.content, nil
}

// IndexOf returns the position of the t
func (l *LinkedList) IndexOf(t T) int {
	l.RLock()
	defer l.RUnlock()

	head := l.head
	idx := 0
	for {
		if head.content == t {
			return idx
		}
		if head.next == nil {
			return -1
		}
		head = head.next
		idx++
	}
}

// IsEmpty returns true if the list is empty
func (l *LinkedList) IsEmpty() bool {
	l.RLock()
	defer l.RUnlock()

	return l.head == nil
}

// Size returns the linked list size
func (l *LinkedList) Size() int {
	l.RLock()
	defer l.RUnlock()

	return l.size
}

// Traverse traverses linked list
func (l *LinkedList) Traverse() {
	l.RLock()
	defer l.RUnlock()

	head := l.head
	for {
		if head == nil {
			break
		}

		head.Print()
		head = head.next
	}
	fmt.Println()
}
