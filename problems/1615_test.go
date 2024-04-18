package problems

import (
	"reflect"
	"testing"
)

func TestMaximalNetworkRank(t *testing.T) {
	testCases := []struct {
		road  int
		roads [][]int
		want  int
	}{
		{4, [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}}, 4},
		{5, [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}, {2, 3}, {2, 4}}, 5},
		{8, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {5, 6}, {5, 7}}, 5},
	}

	for _, tc := range testCases {
		got := maximalNetworkRank(tc.road, tc.roads)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("maximalNetworkRank(%d, %v) = %d; want %d", tc.road, tc.roads, got, tc.want)
		}
	}
}
