package mathematical

import "math"

// EuclideanAlgorithm is finding greatest common divisor (GCD).
func EuclideanAlgorithma(a, b int) int {
	c, d := math.Abs(float64(a)), math.Abs(float64(b))

	if d == 0 {
		return int(c)
	}

	return EuclideanAlgorithma(int(d), int(c)%int(d))
}
