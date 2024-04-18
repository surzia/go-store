package problems

func myPow(x float64, n int) float64 {
	var df func(e float64, p int) float64
	df = func(e float64, p int) float64 {
		if p == 0 {
			return 1
		}
		if p == 1 {
			return e
		}

		var half = df(e, p/2)
		if p%2 == 0 {
			return half * half
		}

		return e * half * half
	}
	if n < 0 {
		return 1 / df(x, -n)
	}

	return df(x, n)
}
