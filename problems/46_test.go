package problems

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{[]int{0, 1}, [][]int{{0, 1}, {1, 0}}},
		{[]int{1}, [][]int{{1}}},
	}

	for _, tc := range testCases {
		result := permute(tc.nums)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("permute(%v) = %v; want %v", tc.nums, result, tc.expected)
		}
	}
}
