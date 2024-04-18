package problems

import "strings"

func removeKdigits(num string, k int) string {
	var stack []byte
	for i := 0; i < len(num); i++ {
		var value = num[i]
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > value {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, value)
	}

	for k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}
	result := strings.TrimLeft(string(stack), "0")
	if result == "" {
		return "0"
	}
	return result
}
