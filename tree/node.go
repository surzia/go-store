package tree

type T int

type Tree struct {
	val   T
	left  *Tree
	right *Tree
}

func NewTree(v T) *Tree {
	tree := &Tree{val: v}

	return tree
}
