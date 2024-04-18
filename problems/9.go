package problems

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	} else {
		var a = make([]int, 0)
		var c = x
		for x > 0 {
			a = append(a, x%10)
			x /= 10
		}
		var e int
		for i := 0; i < len(a); i++ {
			var p = 1
			for j := 0; j < i; j++ {
				p *= 10
			}
			e += a[len(a)-1-i] * p
		}
		return e == c
	}
}
