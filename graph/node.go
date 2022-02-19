package graph

import (
	"fmt"
	"sync"
)

type T string

type Node struct {
	value T
}

type NodeQueue struct {
	sync.RWMutex
	array []*Node
}

func (q *NodeQueue) Enqueue(node *Node) {
	q.Lock()
	q.array = append(q.array, node)
	q.Unlock()
}

func (q *NodeQueue) IsEmpty() bool {
	q.RLock()
	defer q.RUnlock()

	return len(q.array) == 0
}

func (q *NodeQueue) Dequeue() (*Node, error) {
	q.Lock()

	if len(q.array) == 0 {
		q.Unlock()
		return nil, fmt.Errorf("queue is empty")
	}
	ret := q.array[0]
	q.array = q.array[1:len(q.array)]
	q.Unlock()

	return ret, nil
}

// Print a node
func (n *Node) Print() string {
	return fmt.Sprintf("%v", n.value)
}

func NewNodeQueue() *NodeQueue {
	return &NodeQueue{array: make([]*Node, 0)}
}
