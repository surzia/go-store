package problems

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
		// Additional test cases can be added here.
	}

	for _, tc := range testCases {
		nextPermutation(tc.nums)
		if !reflect.DeepEqual(tc.nums, tc.want) {
			t.Errorf("nextPermutation(%v) = %v; want %v", tc.nums, tc.nums, tc.want)
		}
	}
}
