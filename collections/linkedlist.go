package collections

import "errors"

type Node struct {
	value interface{}
	next  *Node
}

type LinkedList struct {
	head *Node
}

// 添加元素到链表尾部
func (l *LinkedList) Append(value interface{}) {
	newNode := &Node{value: value, next: nil}
	if l.head == nil {
		l.head = newNode
		return
	}

	lastNode := l.head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}
	lastNode.next = newNode
}

// 删除特定值的第一个节点
func (l *LinkedList) Delete(value interface{}) error {
	if l.head == nil {
		return errors.New("empty list")
	}

	if l.head.value == value {
		l.head = l.head.next
		return nil
	}

	prev := l.head
	current := l.head.next
	for current != nil {
		if current.value == value {
			prev.next = current.next
			return nil
		}
		prev = current
		current = current.next
	}
	return errors.New("value not found")
}
