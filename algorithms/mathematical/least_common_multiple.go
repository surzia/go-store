package mathematical

import "math"

// LeastCommonMultiple. In arithmetic and number theory, the
// least common multiple, lowest common multiple, or smallest
// common multiple of two integers a and b, usually denoted by
// LeastCommonMultiple(a, b), is the smallest positive integer
// that is divisible by both a and b. Since division of integers
// by zero is undefined, this definition has meaning only if a
// and b are both different from zero. However, some authors
// define LeastCommonMultiple(a,0) as 0 for all a, which is the
// result of taking the lcm to be the least upper bound in the
// lattice of divisibility.
func LeastCommonMultiple(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	return int(math.Abs(float64(a*b))) / EuclideanAlgorithma(a, b)
}
