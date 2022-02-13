package hashmap

// the hash() private function uses the famous Horner's method
// to generate a hash of a string with O(n) complexity
func hash(k Key) int {
	h := 0
	for i := 0; i < len(k); i++ {
		h = 31*h + int(k[i])
	}
	return h
}
