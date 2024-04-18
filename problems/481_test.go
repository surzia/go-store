package problems

import "testing"

func TestMagicalString(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 3},
		{6, 3},
	}

	for _, tc := range testCases {
		result := magicalString(tc.n)
		if result != tc.expected {
			t.Errorf("magicalString(%d) = %d; want %d", tc.n, result, tc.expected)
		}
	}
}
