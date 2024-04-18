package problems

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		strs     []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"abc", "abc", "abc"}, "abc"},
		{[]string{"abc", "ab", "a"}, "a"},
		{[]string{"ab", "a"}, "a"},
		{[]string{"a"}, "a"},
	}

	for _, tc := range testCases {
		result := longestCommonPrefix(tc.strs)
		if result != tc.expected {
			t.Errorf("longestCommonPrefix(%v) = %q; want %q", tc.strs, result, tc.expected)
		}
	}
}
