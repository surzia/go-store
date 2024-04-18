package problems

import "testing"

func TestCircularArrayLoop(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected bool
	}{
		{[]int{2, -1, 1, 2, 2}, true},
		{[]int{-1, 2}, false},
		{[]int{-2, 1, -1, -2, -2}, false},
	}

	for _, tc := range testCases {
		result := circularArrayLoop(tc.nums)
		if result != tc.expected {
			t.Errorf("circularArrayLoop(%v) = %t; want %t", tc.nums, result, tc.expected)
		}
	}
}
