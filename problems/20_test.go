package problems

import "testing"

func TestIsValid(t *testing.T) {
	testCases := []struct {
		s        string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true},
	}

	for _, tc := range testCases {
		result := isValid(tc.s)
		if result != tc.expected {
			t.Errorf("isValid(%s) = %t; want %t", tc.s, result, tc.expected)
		}
	}
}
