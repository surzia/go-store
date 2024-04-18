package problems

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 0}, 3},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 11, 12}, 1},
	}

	for _, tc := range testCases {
		result := firstMissingPositive(tc.nums)
		if result != tc.expected {
			t.Errorf("firstMissingPositive(%v) = %d; want %d", tc.nums, result, tc.expected)
		}
	}
}
