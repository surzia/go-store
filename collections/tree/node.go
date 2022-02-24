package tree

type T int

type Node struct {
	val   T
	left  *Node
	right *Node
}

func NewTree(v T) *Node {
	tree := &Node{val: v}

	return tree
}
