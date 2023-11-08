package collections

import "testing"

func TestGraph_AddNode(t *testing.T) {
	g := NewGraph()

	g.AddNode("A")
	g.AddNode("B")

	_, foundA := g.Nodes["A"]
	_, foundB := g.Nodes["B"]

	if !foundA || !foundB {
		t.Errorf("Expected nodes A and B to be added, but one or both are missing")
	}
}

func TestGraph_AddEdge(t *testing.T) {
	g := NewGraph()

	g.AddNode("A")
	g.AddNode("B")
	g.AddEdge("A", "B")

	if !g.HasEdge("A", "B") || !g.HasEdge("B", "A") {
		t.Error("Expected edge between A and B, but not found")
	}
}

func TestGraph_HasEdge(t *testing.T) {
	g := NewGraph()

	g.AddNode("A")
	g.AddNode("B")
	g.AddEdge("A", "B")

	if !g.HasEdge("A", "B") {
		t.Error("Expected edge between A and B, but not found")
	}

	if g.HasEdge("A", "C") || g.HasEdge("B", "C") {
		t.Error("Expected no edge between A or B and C, but found one")
	}
}
