package graph

import (
	"fmt"
	"testing"
)

var (
	nA = &Node{"A"}
	nB = &Node{"B"}
	nC = &Node{"C"}
	nD = &Node{"D"}
	nE = &Node{"E"}
	nF = &Node{"F"}
)

func InitGraph() *Graph {
	g := NewGraph()
	g.AddNode(nA)
	g.AddNode(nB)
	g.AddNode(nC)
	g.AddNode(nD)
	g.AddNode(nE)
	g.AddNode(nF)

	g.AddEdge(nA, nB)
	g.AddEdge(nA, nC)
	g.AddEdge(nB, nE)
	g.AddEdge(nC, nE)
	g.AddEdge(nE, nF)
	g.AddEdge(nD, nA)

	return g
}

func TestGraph_Print(t *testing.T) {
	g := InitGraph()
	g.Print()
}

func TestGraph_Traverse(t *testing.T) {
	g := InitGraph()
	g.Traverse(func(n *Node) {
		fmt.Printf("%v\n", n.value)
	})
}
