package sort

import (
	"math/rand"
	"testing"
)

func InitUnsortedArray(n int) ([]int, []int) {
	old := make([]int, n)
	for i := 0; i < n; i++ {
		old[i] = i
	}
	ret := make([]int, n)
	perm := rand.Perm(n)
	for i := 0; i < n; i++ {
		ret[i] = old[perm[i]]
	}

	return ret, old
}

func TestBubbleSort(t *testing.T) {
	arr, sorted := InitUnsortedArray(10)
	BubbleSort(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] != sorted[i] {
			t.Errorf("index %d of arr should be %d after sorting, but got %d", i, sorted[i], arr[i])
		}
	}
}

func TestOptimizedBubbleSort(t *testing.T) {
	arr, sorted := InitUnsortedArray(10)
	OptimizedBubbleSort(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] != sorted[i] {
			t.Errorf("index %d of arr should be %d after sorting, but got %d", i, sorted[i], arr[i])
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	arr, _ := InitUnsortedArray(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(arr)
	}
}

func BenchmarkOptimizedBubbleSort(b *testing.B) {
	arr, _ := InitUnsortedArray(10)
	for i := 0; i < b.N; i++ {
		OptimizedBubbleSort(arr)
	}
}

func TestQuickSort(t *testing.T) {
	arr, sorted := InitUnsortedArray(10)
	QuickSort(arr, 0, len(arr)-1)
	for i := 0; i < len(arr); i++ {
		if arr[i] != sorted[i] {
			t.Errorf("index %d of arr should be %d after sorting, but got %d", i, sorted[i], arr[i])
		}
	}
}
