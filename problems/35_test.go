package problems

import "testing"

func TestSearchInsert(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{1, 3, 5}, 3, 1},
		{[]int{1, 3, 5}, 4, 2},
		{[]int{1, 3, 5}, 2, 1},
		{[]int{2, 4}, 2, 0},
		{[]int{2, 4}, 4, 1},
		{[]int{2, 4}, 3, 1},
	}

	for _, tc := range testCases {
		result := searchInsert(tc.nums, tc.target)
		if result != tc.expected {
			t.Errorf("searchInsert(%v, %d) = %d; want %d", tc.nums, tc.target, result, tc.expected)
		}
	}
}
