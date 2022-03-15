package search

import (
	"testing"
)

func InitSortedArray(n int) []int {
	ret := make([]int, n)

	for i := 0; i < n; i++ {
		ret[i] = i
	}
	return ret
}

func TestIterativeBinarySearch(t *testing.T) {
	array := InitSortedArray(6)
	index := IterativeBinarySearch(array, 5, 0, len(array)-1)
	if index != 5 {
		t.Errorf("index should be %d, but got %d", 5, index)
	}

	index = IterativeBinarySearch(array, 7, 0, len(array)-1)
	if index != -1 {
		t.Errorf("index should be %d, but got %d", -1, index)
	}
}

func TestRecursiveBinarySearch(t *testing.T) {
	array := InitSortedArray(8)
	index := RecursiveBinarySearch(array, 5, 0, len(array)-1)
	if index != 5 {
		t.Errorf("index should be %d, but got %d", 5, index)
	}

	index = IterativeBinarySearch(array, 9, 0, len(array)-1)
	if index != -1 {
		t.Errorf("index should be %d, but got %d", -1, index)
	}
}

func BenchmarkIterativeBinarySearch(b *testing.B) {
	array := InitSortedArray(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IterativeBinarySearch(array, 678, 0, len(array)-1)
	}
}

func BenchmarkRecursiveBinarySearch(b *testing.B) {
	array := InitSortedArray(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RecursiveBinarySearch(array, 678, 0, len(array)-1)
	}
}
