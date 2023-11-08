package algorithms

import "testing"

func TestLinearSearch(t *testing.T) {
	arr := []int{11, 25, 12, 64, 22}
	target := 12
	expected := 2

	result := LinearSearch(arr, target)

	if result != expected {
		t.Errorf("LinearSearch: Expected index %d, got %d", expected, result)
	}
}

func TestBinarySearch(t *testing.T) {
	arr := []int{11, 12, 22, 25, 64}
	target := 22
	expected := 2

	result := BinarySearch(arr, target)

	if result != expected {
		t.Errorf("BinarySearch: Expected index %d, got %d", expected, result)
	}
}

func TestBFSSearch(t *testing.T) {
	graph := map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {6},
		4: {},
		5: {},
		6: {7},
		7: {},
	}
	startNode := 1
	targetNode := 7
	expected := true

	result := BFS(graph, startNode, targetNode)

	if result != expected {
		t.Errorf("BFSSearch: Expected path between nodes, got %t", result)
	}
}

func TestDFSSearch(t *testing.T) {
	graph := map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {6},
		4: {},
		5: {},
		6: {7},
		7: {},
	}
	startNode := 1
	targetNode := 7
	expected := true

	result := DFS(graph, startNode, targetNode)

	if result != expected {
		t.Errorf("DFSSearch: Expected path between nodes, got %t", result)
	}
}
