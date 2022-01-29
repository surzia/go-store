package linkedlist

import "fmt"

type T int

// Node a single node that composes the list
type Node struct {
	content T
	next    *Node
}

func NewNode(t T) *Node {
	return &Node{content: t}
}

func (n *Node) Print() {
	fmt.Print(n.content)
	fmt.Print(" ")
}
