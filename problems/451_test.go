package problems

import (
	"reflect"
	"testing"
)

func TestFrequencySort(t *testing.T) {
	testCases := []struct {
		s        string
		expected string
	}{
		{"tree", "eetr"},
		{"cccaaa", "cccaaa"},
		{"Aabb", "bbAa"},
	}

	for _, tc := range testCases {
		result := frequencySort(tc.s)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("frequencySort(%s) = %s; want %s", tc.s, result, tc.expected)
		}
	}
}
