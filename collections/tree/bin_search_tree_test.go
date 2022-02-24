package tree

import (
	"testing"
)

const (
	t1 T = iota + 1
	t2
	t3
	t4
	t5
	t6
	t7
	t8
	t9
	t10
	t11
)

func InitTree() *BinSearchTree {
	bst := NewBinSearchTree()
	bst.Insert(t3)
	bst.Insert(t2)
	bst.Insert(t4)
	bst.Insert(t9)
	bst.Insert(t5)
	bst.Insert(t6)
	bst.Insert(t10)
	bst.Insert(t1)
	bst.Insert(t7)
	bst.Insert(t8)

	return bst
}

func TestBinSearchTree_Insert(t *testing.T) {
	bst := InitTree()
	bst.Print()

	bst.Insert(t11)
	bst.Print()
}

func TestBinSearchTree_InOrder(t *testing.T) {
	var ret []T
	bst := InitTree()
	bst.InOrder(func(t T) {
		ret = append(ret, t)
	})

	expected := []T{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !isSameSlice(ret, expected) {
		t.Errorf("Traversal order incorrect, got %v", ret)
	}
}

// isSameSlice returns true if the 2 slices are identical
func isSameSlice(a, b []T) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
