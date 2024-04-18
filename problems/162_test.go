package problems

import "testing"

func TestFindPeakElement(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 1}, 2},
		{[]int{1, 2, 1, 3, 5, 6, 4}, 5},
		{[]int{1, 2}, 1},
		{[]int{3, 2}, 0},
	}

	for _, tc := range testCases {
		result := findPeakElement(tc.nums)
		if result != tc.expected {
			t.Errorf("findPeakElement(%v) = %d; want %d", tc.nums, result, tc.expected)
		}
	}
}
