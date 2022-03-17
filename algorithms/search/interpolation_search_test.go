package search

import "testing"

func InitArray() []int {
	ret := []int{13, 21, 34, 55, 69, 73, 84, 101}
	return ret
}

func TestInterpolationSearch(t *testing.T) {
	arr := InitArray()
	ret := InterpolationSearch(arr, 84)
	if ret != 6 {
		t.Errorf("expected index %d, but got %d", 6, ret)
	}

	ret = InterpolationSearch(arr, 19)
	if ret != -1 {
		t.Errorf("expected index %d, but got %d", -1, ret)
	}
}
