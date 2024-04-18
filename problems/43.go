package problems

import "strconv"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	if len(num1) > len(num2) {
		return multiply(num2, num1)
	}

	var df func(s1, s2 string) string
	df = func(s1, s2 string) string {
		if s1 == "" {
			return s2
		}
		if s2 == "" {
			return s1
		}

		var point = 0
		var ans = ""

		var carry = 0
		for point < len(s1) || point < len(s2) {
			var left, right = 0, 0
			if point < len(s1) {
				left = int(s1[point] - '0')
			}
			if point < len(s2) {
				right = int(s2[point] - '0')
			}
			var res = left + right + carry
			carry = res / 10
			ans += strconv.Itoa(res % 10)
			point++
		}
		if carry != 0 {
			ans += strconv.Itoa(carry)
		}

		return ans
	}

	var reverseString func(s string) string
	reverseString = func(s string) string {
		runes := []rune(s)

		left, right := 0, len(runes)-1

		for left < right {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}

		return string(runes)
	}

	var res = ""
	for i := len(num1) - 1; i >= 0; i-- {
		var tmp = ""
		var k = num2
		for j := len(num1) - 1; j > i; j-- {
			k += "0"
		}
		for j := 0; j < int(num1[i]-'0'); j++ {
			tmp = df(tmp, reverseString(k))
		}
		res = df(res, tmp)
	}
	return reverseString(res)
}
