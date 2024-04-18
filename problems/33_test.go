package problems

import "testing"

func TestSearch(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{1, 3}, 3, 1},
	}

	for _, tc := range testCases {
		result := search(tc.nums, tc.target)
		if result != tc.expected {
			t.Errorf("search(%v, %d) = %d; want %d", tc.nums, tc.target, result, tc.expected)
		}
	}
}
