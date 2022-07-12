package mathematical

import (
	"errors"
	"math"
)

const topMaxValidPosition = 70

// Fibonacci returns a sequence, and characterized by the fact
// that every number after the first two is the sum of the two
// preceding ones: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ...
func Fibonacci(n int) []int {
	if n < 1 {
		panic(errors.New("n cannot be smaller than 1"))
	}
	var fibSequence = []int{1}

	currentValue, previousValue := 1, 0
	if n == 1 {
		return fibSequence
	}

	iterationsCounter := n - 1
	for {
		if iterationsCounter <= 0 {
			break
		}

		currentValue += previousValue
		previousValue += currentValue
		fibSequence = append(fibSequence, currentValue)
		iterationsCounter -= 1
	}

	return fibSequence
}

// FibonacciNth returns n-th number in Fibonacci sequence
func FibonacciNth(n int) int {
	if n < 1 {
		panic(errors.New("n cannot be smaller than 1"))
	}

	if n == 1 {
		return 1
	}

	currentValue, previousValue, iterationsCounter := 1, 0, n-1
	for {
		if iterationsCounter <= 0 {
			break
		}

		currentValue += previousValue
		previousValue += currentValue
		iterationsCounter -= 1
	}

	return currentValue
}

// FibonacciNthClosedForm returns n-th number in Fibonacci sequence
// according to https://en.wikipedia.org/wiki/Fibonacci_number#Closed-form_expression
func FibonacciNthClosedForm(n int) int {
	if n < 1 || n > topMaxValidPosition {
		panic(errors.New("n cannot be smaller than 1 or larger than 70"))
	}

	// calculate √5 to re-use it in further formulas
	sqrt5 := math.Sqrt(5)
	// calculate φ constant (≈ 1.61803)
	phi := (1 + sqrt5) / 2

	return int(math.Floor(math.Pow(phi, float64(n))/sqrt5 + 0.5))
}
