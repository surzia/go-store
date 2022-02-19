package graph

import "fmt"

type T string

type Node struct {
	value T
}

// Print a node
func (n *Node) Print() string {
	return fmt.Sprintf("%v", n.value)
}
