package algorithms

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{64, 25, 12, 22, 11}
	expected := []int{11, 12, 22, 25, 64}

	BubbleSort(arr)

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("BubbleSort: Expected %v, got %v", expected, arr)
	}
}

func TestInsertionSort(t *testing.T) {
	arr := []int{64, 25, 12, 22, 11}
	expected := []int{11, 12, 22, 25, 64}

	InsertionSort(arr)

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("InsertionSort: Expected %v, got %v", expected, arr)
	}
}

func TestSelectionSort(t *testing.T) {
	arr := []int{64, 25, 12, 22, 11}
	expected := []int{11, 12, 22, 25, 64}

	SelectionSort(arr)

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("SelectionSort: Expected %v, got %v", expected, arr)
	}
}

func TestMergeSort(t *testing.T) {
	arr := []int{64, 25, 12, 22, 11}
	expected := []int{11, 12, 22, 25, 64}

	MergeSort(arr)

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("MergeSort: Expected %v, got %v", expected, arr)
	}
}

func TestQuickSort(t *testing.T) {
	arr := []int{64, 25, 12, 22, 11}
	expected := []int{11, 12, 22, 25, 64}

	QuickSort(arr)

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("QuickSort: Expected %v, got %v", expected, arr)
	}
}

func TestHeapSort(t *testing.T) {
	arr := []int{64, 25, 12, 22, 11}
	expected := []int{11, 12, 22, 25, 64}

	HeapSort(arr)

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("HeapSort: Expected %v, got %v", expected, arr)
	}
}

func TestCountingSort(t *testing.T) {
	arr := []int{64, 25, 12, 22, 11}
	expected := []int{11, 12, 22, 25, 64}

	CountingSort(arr)

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("CountingSort: Expected %v, got %v", expected, arr)
	}
}

func TestBucketSort(t *testing.T) {
	arr := []int{64, 25, 12, 22, 11}
	expected := []int{11, 12, 22, 25, 64}

	BucketSort(arr)

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("BucketSort: Expected %v, got %v", expected, arr)
	}
}

func TestRadixSort(t *testing.T) {
	arr := []int{64, 25, 12, 22, 11}
	expected := []int{11, 12, 22, 25, 64}

	RadixSort(arr)

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("RadixSort: Expected %v, got %v", expected, arr)
	}
}
