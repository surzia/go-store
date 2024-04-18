package problems

import (
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	testCases := []struct {
		x        float64
		n        int
		expected float64
	}{
		{2.00000, 10, 1024.00000},
		{2.10000, 3, 9.26100},
		{2.00000, -2, 0.25000},
	}

	epsilon := 1e-5

	for _, tc := range testCases {
		result := myPow(tc.x, tc.n)
		if math.Abs(result-tc.expected) > epsilon {
			t.Errorf("myPow(%f, %d) = %f; want %f", tc.x, tc.n, result, tc.expected)
		}
	}
}
