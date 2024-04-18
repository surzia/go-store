package problems

import "testing"

func TestCountAndSay(t *testing.T) {
	testCases := []struct {
		n        int
		expected string
	}{
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
		{5, "111221"},
		{6, "312211"},
	}

	for _, tc := range testCases {
		result := countAndSay(tc.n)
		if result != tc.expected {
			t.Errorf("countAndSay(%d) = %s; want %s", tc.n, result, tc.expected)
		}
	}
}
