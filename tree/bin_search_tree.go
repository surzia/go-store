package tree

import (
	"fmt"
	"sync"
)

type BinSearchTree struct {
	sync.RWMutex
	root *Node
}

func NewBinSearchTree() *BinSearchTree {
	bst := &BinSearchTree{}

	return bst
}

// Insert inserts the val in the tree
func (bst *BinSearchTree) Insert(val T) {
	bst.Lock()
	defer bst.Unlock()

	node := NewTree(val)
	if bst.root == nil {
		bst.root = node
	} else {
		insert(bst.root, node)
	}
}

// insert internal function to find the correct place for a node in a tree
func insert(root, node *Node) {
	if node.val < root.val {
		if root.left == nil {
			root.left = node
		} else {
			insert(root.left, node)
		}
	} else {
		if root.right == nil {
			root.right = node
		} else {
			insert(root.right, node)
		}
	}
}

// InOrder visits all nodes with in-order traversing
func (bst *BinSearchTree) InOrder(f func(T)) {
	bst.Lock()
	defer bst.Unlock()

	inOrder(bst.root, f)
}

// inOrder internal recursive function to traverse in-order
func inOrder(root *Node, f func(T)) {
	if root != nil {
		inOrder(root.left, f)
		f(root.val)
		inOrder(root.right, f)
	}
}

// PreOrder visits all nodes with pre-order traversing
func (bst *BinSearchTree) PreOrder(f func(T)) {
	bst.Lock()
	defer bst.Unlock()

	preOrder(bst.root, f)
}

// preOrder internal recursive function to traverse pre-order
func preOrder(root *Node, f func(T)) {
	if root != nil {
		f(root.val)
		inOrder(root.left, f)
		inOrder(root.right, f)
	}
}

// PostOrder visits all nodes with post-order traversing
func (bst *BinSearchTree) PostOrder(f func(T)) {
	bst.Lock()
	defer bst.Unlock()

	postOrder(bst.root, f)
}

// postOrder internal recursive function to traverse post-order
func postOrder(root *Node, f func(T)) {
	if root != nil {
		inOrder(root.left, f)
		inOrder(root.right, f)
		f(root.val)
	}
}

// Min returns the minimal value stored in the tree
func (bst *BinSearchTree) Min() *T {
	bst.RLock()
	defer bst.RUnlock()

	root := bst.root
	if root == nil {
		return nil
	}
	for {
		if root.left == nil {
			return &root.val
		}
		root = root.left
	}
}

// Max returns the maximal value stored in the tree
func (bst *BinSearchTree) Max() *T {
	bst.RLock()
	defer bst.RUnlock()

	root := bst.root
	if root == nil {
		return nil
	}
	for {
		if root.right == nil {
			return &root.val
		}
		root = root.right
	}
}

// Search returns true if the t exists in the tree
func (bst *BinSearchTree) Search(t T) bool {
	bst.RLock()
	defer bst.RUnlock()

	return search(bst.root, t)
}

// search internal recursive function to search t in the tree
func search(root *Node, t T) bool {
	if root == nil {
		return false
	}

	if root.val == t {
		return true
	}

	if root.val > t {
		return search(root.left, t)
	} else {
		return search(root.right, t)
	}
}

// Remove removes the t from the tree
func (bst *BinSearchTree) Remove(t T) {
	bst.Lock()
	defer bst.Unlock()

	remove(bst.root, t)
}

// remove internal recursive function to remove t
func remove(root *Node, t T) {
	if root == nil {
		return
	}

	if root.val > t {
		remove(root.left, t)
	} else if root.val < t {
		remove(root.right, t)
	} else {
		if root.left == nil && root.right == nil {
			root = nil
			return
		} else if root.left == nil {
			root = root.right
			return
		} else if root.right == nil {
			root = root.left
			return
		} else {
			leftMostRightSide := root.right
			for {
				//find the smallest value on the right side
				if leftMostRightSide != nil && leftMostRightSide.left != nil {
					leftMostRightSide = leftMostRightSide.left
				} else {
					break
				}
			}

			root.val = leftMostRightSide.val
			remove(root.right, root.val)
			return
		}
	}
}

func (bst *BinSearchTree) Print() {
	bst.Lock()
	defer bst.Unlock()

	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

func stringify(root *Node, level int) {
	if root != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(root.left, level)
		fmt.Printf(format+"%d\n", root.val)
		stringify(root.right, level)
	}
}
