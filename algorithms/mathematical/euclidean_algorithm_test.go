package mathematical

import "testing"

func TestEuclideanAlgorithma(t *testing.T) {
	a, b, c := 0, 0, 0
	if c != EuclideanAlgorithma(a, b) {
		t.Errorf("expect %d, but got %d", c, EuclideanAlgorithma(a, b))
	}

	a, b, c = 2, 0, 2
	if c != EuclideanAlgorithma(a, b) {
		t.Errorf("expect %d, but got %d", c, EuclideanAlgorithma(a, b))
	}

	a, b, c = 0, 2, 2
	if c != EuclideanAlgorithma(a, b) {
		t.Errorf("expect %d, but got %d", c, EuclideanAlgorithma(a, b))
	}

	a, b, c = 1, 2, 1
	if c != EuclideanAlgorithma(a, b) {
		t.Errorf("expect %d, but got %d", c, EuclideanAlgorithma(a, b))
	}

	a, b, c = 2, 1, 1
	if c != EuclideanAlgorithma(a, b) {
		t.Errorf("expect %d, but got %d", c, EuclideanAlgorithma(a, b))
	}

	a, b, c = 6, 6, 6
	if c != EuclideanAlgorithma(a, b) {
		t.Errorf("expect %d, but got %d", c, EuclideanAlgorithma(a, b))
	}

	a, b, c = 2, 4, 2
	if c != EuclideanAlgorithma(a, b) {
		t.Errorf("expect %d, but got %d", c, EuclideanAlgorithma(a, b))
	}

	a, b, c = 4, 2, 2
	if c != EuclideanAlgorithma(a, b) {
		t.Errorf("expect %d, but got %d", c, EuclideanAlgorithma(a, b))
	}

	a, b, c = 12, 4, 4
	if c != EuclideanAlgorithma(a, b) {
		t.Errorf("expect %d, but got %d", c, EuclideanAlgorithma(a, b))
	}
}
