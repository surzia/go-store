package problems

import "testing"

func TestCanJump(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{0}, true},
		{[]int{1, 0, 1, 0}, false},
	}

	for _, tc := range testCases {
		result := canJump(tc.nums)
		if result != tc.expected {
			t.Errorf("canJump(%v) = %t; want %t", tc.nums, result, tc.expected)
		}
	}
}
