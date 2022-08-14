package mathematical

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	n := 4
	ret := IsPowerOfTwo(n)
	if !ret {
		t.Errorf("%d is power of 2", n)
	}

	n = 34
	ret = IsPowerOfTwo(n)
	if ret {
		t.Errorf("%d is not power of 2", n)
	}
}
