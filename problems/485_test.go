package problems

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 1, 0, 1, 1, 1}, 3},
		{[]int{1, 0, 1, 1, 0, 1}, 2},
		{[]int{0, 0, 0, 0}, 0},
	}

	for _, tc := range testCases {
		result := findMaxConsecutiveOnes(tc.nums)
		if result != tc.expected {
			t.Errorf("findMaxConsecutiveOnes(%v) = %d; want %d", tc.nums, result, tc.expected)
		}
	}
}
