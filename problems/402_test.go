package problems

import "testing"

func TestRemoveKdigits(t *testing.T) {
	testCases := []struct {
		num      string
		k        int
		expected string
	}{
		{"1432219", 3, "1219"},
		{"10200", 1, "200"},
		{"10", 2, "0"},
		{"9", 1, "0"},
	}

	for _, tc := range testCases {
		result := removeKdigits(tc.num, tc.k)
		if result != tc.expected {
			t.Errorf("removeKdigits(%s, %d) = %s; want %s", tc.num, tc.k, result, tc.expected)
		}
	}
}
