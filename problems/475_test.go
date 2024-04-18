package problems

import "testing"

func TestFindRadius(t *testing.T) {
	testCases := []struct {
		houses   []int
		heaters  []int
		expected int
	}{
		{[]int{1, 2, 3}, []int{2}, 1},
		{[]int{1, 2, 3, 4}, []int{1, 4}, 1},
		{[]int{1, 5}, []int{2}, 3},
	}

	for _, tc := range testCases {
		result := findRadius(tc.houses, tc.heaters)
		if result != tc.expected {
			t.Errorf("findRadius(%v, %v) = %d; want %d", tc.houses, tc.heaters, result, tc.expected)
		}
	}
}
