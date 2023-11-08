package collections

import "testing"

func TestBinaryTree_Insert(t *testing.T) {
	bTree := BinaryTree{}

	bTree.Insert(5)
	bTree.Insert(3)
	bTree.Insert(8)

	if bTree.Root.Value != 5 {
		t.Errorf("Expected root value: 5, got: %d", bTree.Root.Value)
	}
	if bTree.Root.Left.Value != 3 {
		t.Errorf("Expected left child value: 3, got: %d", bTree.Root.Left.Value)
	}
	if bTree.Root.Right.Value != 8 {
		t.Errorf("Expected right child value: 8, got: %d", bTree.Root.Right.Value)
	}
}
