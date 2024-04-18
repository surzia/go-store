package problems

import "testing"

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		height   []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
	}

	for _, tc := range testCases {
		result := maxArea(tc.height)
		if result != tc.expected {
			t.Errorf("maxArea(%v) = %d; want %d", tc.height, result, tc.expected)
		}
	}
}
