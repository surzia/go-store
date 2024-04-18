package problems

import "testing"

func TestMinMoves2(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3}, 2},
		{[]int{1, 10, 2, 9}, 16},
		{[]int{1}, 0},
		{[]int{1, 1}, 0},
	}

	for _, tc := range testCases {
		result := minMoves2(tc.nums)
		if result != tc.expected {
			t.Errorf("minMoves2(%v) = %d; want %d", tc.nums, result, tc.expected)
		}
	}
}
