package problems

func generateParenthesis(n int) []string {
	var res []string
	var df func(left, right int, p string)
	df = func(left, right int, p string) {
		if left == n && right == n {
			res = append(res, p)
			return
		}

		if left < right || left > n || right > n {
			return
		}

		df(left+1, right, p+"(")
		df(left, right+1, p+")")
	}

	df(0, 0, "")
	return res
}
