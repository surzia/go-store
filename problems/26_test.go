package problems

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, tc := range testCases {
		result := removeDuplicates(tc.nums)
		if result != tc.expected {
			t.Errorf("removeDuplicates(%v) = %d; want %d", tc.nums, result, tc.expected)
		}
	}
}
