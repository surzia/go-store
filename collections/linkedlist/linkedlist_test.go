package linkedlist

import "testing"

var (
	t1 T = 1
	t2 T = 2
	t3 T = 3
	t4 T = 4
)

func InitLinkedList() *LinkedList {
	list := NewLinkedList()
	list.head = NewNode(t1)
	list.head.next = NewNode(t2)
	list.head.next.next = NewNode(t3)
	list.size = 3

	return list
}

func TestLinkedList_Add(t *testing.T) {
	linkedList := InitLinkedList()
	size := linkedList.Size()

	if size != 3 {
		t.Errorf("linked list size should be 3 but got %d", size)
	}

	linkedList.Add(4)
	size = linkedList.Size()
	if size != 4 {
		t.Errorf("linked list size should be 4 but got %d", size)
	}
}

func TestLinkedList_Insert(t *testing.T) {
	linkedList := InitLinkedList()

	err := linkedList.Insert(t4, 1)
	if err != nil {
		t.Errorf("insert into linked list err %v", err)
	}

	idx1 := linkedList.IndexOf(t2)
	idx2 := linkedList.IndexOf(t4)
	if idx1 != 2 || idx2 != 1 {
		t.Errorf("linked list position is not expected.")
	}
}

func TestLinkedList_RemoveAt(t *testing.T) {
	linkedList := InitLinkedList()

	ret, err := linkedList.RemoveAt(2)
	if err != nil {
		t.Errorf("linked list err %v", err)
	}
	if *ret != t3 {
		t.Errorf("removed item expect 3 but got %d", *ret)
	}

	size := linkedList.Size()
	if size != 2 {
		t.Errorf("linked list size should be 2 but got %d", size)
	}
}

func TestLinkedList_IsEmpty(t *testing.T) {
	linkedList := InitLinkedList()

	empty := linkedList.IsEmpty()
	if empty {
		t.Errorf("linked list is not empty.")
	}

	linkedList = &LinkedList{}
	empty = linkedList.IsEmpty()
	if !empty {
		t.Errorf("linked list is empty.")
	}
}

func TestLinkedList_Traverse(t *testing.T) {
	linkedList := InitLinkedList()

	linkedList.Traverse()
}
