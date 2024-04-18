package problems

import "testing"

func TestCheckRecord(t *testing.T) {
	testCases := []struct {
		s        string
		expected bool
	}{
		{"PPALLP", true},
		{"PPALLL", false},
		{"AAAA", false},
		{"LLLL", false},
		{"", true},
	}

	for _, tc := range testCases {
		result := checkRecord(tc.s)
		if result != tc.expected {
			t.Errorf("checkRecord(%s) = %t; want %t", tc.s, result, tc.expected)
		}
	}
}
