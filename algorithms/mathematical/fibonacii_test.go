package mathematical

import "testing"

func TestFibonacci(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("the code did not panic")
		}
	}()
	Fibonacci(0)

	n := 8
	expect := []int{1, 1, 2, 3, 5, 8, 13, 21}
	real := Fibonacci(n)
	for i := 0; i < n; i++ {
		if expect[i] != real[i] {
			t.Errorf("sequence i-th index should be %d, but got %d", expect[i], real[i])
		}
	}
}

func TestFibonacciNth(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("the code did not panic")
		}
	}()
	FibonacciNth(0)

	expect, n := 1, 1
	real := FibonacciNth(n)
	if expect != real {
		t.Errorf("sequence i-th index should be %d, but got %d", expect, real)
	}

	expect, n = 2, 1
	real = FibonacciNth(n)
	if expect != real {
		t.Errorf("sequence i-th index should be %d, but got %d", expect, real)
	}

	expect, n = 8, 21
	real = FibonacciNth(n)
	if expect != real {
		t.Errorf("sequence i-th index should be %d, but got %d", expect, real)
	}
}

func TestFibonacciNthClosedForm(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("the code did not panic")
		}
	}()
	FibonacciNthClosedForm(0)

	expect, n := 1, 1
	real := FibonacciNthClosedForm(n)
	if expect != real {
		t.Errorf("sequence i-th index should be %d, but got %d", expect, real)
	}

	expect, n = 2, 1
	real = FibonacciNthClosedForm(n)
	if expect != real {
		t.Errorf("sequence i-th index should be %d, but got %d", expect, real)
	}

	expect, n = 8, 21
	real = FibonacciNthClosedForm(n)
	if expect != real {
		t.Errorf("sequence i-th index should be %d, but got %d", expect, real)
	}
}
