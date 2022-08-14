package mathematical

// IsPowerOfTwo Given a positive integer to find
// if it is a power of two or not.
func IsPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}

	return (n & (n - 1)) == 0
}
