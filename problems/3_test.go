package problems

import (
	"reflect"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		str  string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for _, tc := range testCases {
		got := lengthOfLongestSubstring(tc.str)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("lengthOfLongestSubstring(%s) = %d; want %d", tc.str, got, tc.want)
		}
	}

}
