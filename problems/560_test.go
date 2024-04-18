package problems

import "testing"

func TestSubarraySum(t *testing.T) {
	testCases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 2, 3}, 3, 2},
		{[]int{1, 2, 1, 2, 1}, 3, 4},
	}

	for _, tc := range testCases {
		result := subarraySum(tc.nums, tc.k)
		if result != tc.expected {
			t.Errorf("subarraySum(%v, %d) = %d; want %d", tc.nums, tc.k, result, tc.expected)
		}
	}
}
