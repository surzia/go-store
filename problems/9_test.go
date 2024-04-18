package problems

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{-101, false},
		{12321, true},
		{0, true},
	}

	for _, tc := range testCases {
		result := isPalindrome(tc.input)
		if result != tc.expected {
			t.Errorf("isPalindrome(%d) = %t; want %t", tc.input, result, tc.expected)
		}
	}
}
