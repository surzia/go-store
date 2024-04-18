package problems

import "testing"

func TestJump(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{1, 2, 3}, 2},
		{[]int{1, 2, 1, 1, 1}, 3},
		{[]int{1}, 0},
		{[]int{0}, 0},
	}

	for _, tc := range testCases {
		result := jump(tc.nums)
		if result != tc.expected {
			t.Errorf("jump(%v) = %d; want %d", tc.nums, result, tc.expected)
		}
	}
}
