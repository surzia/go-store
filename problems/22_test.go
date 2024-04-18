package problems

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	testCases := []struct {
		n        int
		expected []string
	}{
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{1, []string{"()"}},
		{2, []string{"(())", "()()"}},
	}

	for _, tc := range testCases {
		result := generateParenthesis(tc.n)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("generateParenthesis(%d) = %v; want %v", tc.n, result, tc.expected)
		}
	}
}
