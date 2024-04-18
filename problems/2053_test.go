package problems

import "testing"

func TestKthDistinct(t *testing.T) {
	testCases := []struct {
		arr      []string
		k        int
		expected string
	}{
		{[]string{"a", "b", "c", "b", "a"}, 1, "c"},
		{[]string{"a", "b", "c", "b", "a"}, 2, ""},
		{[]string{"a", "b", "c", "b", "a"}, 3, ""},
		{[]string{"a", "b", "c", "b", "a"}, 4, ""},
		{[]string{"a", "b", "c", "b", "a"}, 5, ""},
		{[]string{"a", "b", "c", "b", "a"}, 6, ""},
	}

	for _, tc := range testCases {
		result := kthDistinct(tc.arr, tc.k)
		if result != tc.expected {
			t.Errorf("kthDistinct(%v, %d) = %s; want %s", tc.arr, tc.k, result, tc.expected)
		}
	}
}
