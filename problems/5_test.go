package problems

import "testing"

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		s        string
		expected string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"ac", "a"},
	}

	for _, tc := range testCases {
		result := longestPalindrome(tc.s)
		if result != tc.expected {
			t.Errorf("longestPalindrome(%s) = %s; want %s", tc.s, result, tc.expected)
		}
	}
}
