package graph

import (
	"fmt"
	"sync"
)

type Graph struct {
	sync.RWMutex
	nodes []*Node
	edges map[Node][]*Node
}

func NewGraph() *Graph {
	g := &Graph{
		nodes: []*Node{},
		edges: make(map[Node][]*Node),
	}

	return g
}

// AddNode adds a node to the graph
func (g *Graph) AddNode(n *Node) {
	g.Lock()
	g.nodes = append(g.nodes, n)
	g.Unlock()
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(n1, n2 *Node) {
	g.Lock()
	defer g.Unlock()

	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
}

// Print whole graph
func (g *Graph) Print() {
	g.Lock()
	defer g.Unlock()

	ret := ""
	for i := 0; i < len(g.nodes); i++ {
		ret += g.nodes[i].Print() + " -> "
		neighborhood := g.edges[*g.nodes[i]]
		for j := 0; j < len(neighborhood); j++ {
			ret += neighborhood[j].Print() + " "
		}
		ret += "\n"
	}

	fmt.Println(ret)
}

// Traverse implements the BFS traversing algorithm
func (g *Graph) Traverse(f func(*Node)) {
	g.RLock()
	q := NewNodeQueue()
	n := g.nodes[0]
	q.Enqueue(n)
	visited := make(map[*Node]bool)
	for {
		if q.IsEmpty() {
			break
		}
		node, _ := q.Dequeue()
		visited[node] = true
		near := g.edges[*node]

		for i := 0; i < len(near); i++ {
			j := near[i]
			if !visited[j] {
				q.Enqueue(j)
				visited[j] = true
			}
		}
		if f != nil {
			f(node)
		}
	}
	g.RUnlock()
}
