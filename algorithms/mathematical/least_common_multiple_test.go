package mathematical

import "testing"

func TestLeastCommonMultiple(t *testing.T) {
	a, b, c := 0, 0, 0
	if c != LeastCommonMultiple(a, b) {
		t.Errorf("expect %d, but got %d", c, LeastCommonMultiple(a, b))
	}

	a, b, c = 1, 0, 0
	if c != LeastCommonMultiple(a, b) {
		t.Errorf("expect %d, but got %d", c, LeastCommonMultiple(a, b))
	}

	a, b, c = 0, 1, 0
	if c != LeastCommonMultiple(a, b) {
		t.Errorf("expect %d, but got %d", c, LeastCommonMultiple(a, b))
	}

	a, b, c = 4, 6, 12
	if c != LeastCommonMultiple(a, b) {
		t.Errorf("expect %d, but got %d", c, LeastCommonMultiple(a, b))
	}

	a, b, c = 6, 21, 42
	if c != LeastCommonMultiple(a, b) {
		t.Errorf("expect %d, but got %d", c, LeastCommonMultiple(a, b))
	}

	a, b, c = 7, 2, 14
	if c != LeastCommonMultiple(a, b) {
		t.Errorf("expect %d, but got %d", c, LeastCommonMultiple(a, b))
	}

	a, b, c = 3, 5, 15
	if c != LeastCommonMultiple(a, b) {
		t.Errorf("expect %d, but got %d", c, LeastCommonMultiple(a, b))
	}

	a, b, c = -9, -18, 18
	if c != LeastCommonMultiple(a, b) {
		t.Errorf("expect %d, but got %d", c, LeastCommonMultiple(a, b))
	}

	a, b, c = 1000000, 2, 1000000
	if c != LeastCommonMultiple(a, b) {
		t.Errorf("expect %d, but got %d", c, LeastCommonMultiple(a, b))
	}
}
