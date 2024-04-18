package problems

import "testing"

func TestConvert(t *testing.T) {
	testCases := []struct {
		s        string
		numRows  int
		expected string
	}{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"AB", 1, "AB"},
		{"A", 1, "A"},
	}

	for _, tc := range testCases {
		result := convert(tc.s, tc.numRows)
		if result != tc.expected {
			t.Errorf("convert(%q, %d) = %q; want %q", tc.s, tc.numRows, result, tc.expected)
		}
	}
}
