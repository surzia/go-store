package problems

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				dp[i][j] = true
			} else {
				dp[i][j] = false
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if s[i] == s[j] {
				if j+1 > i-1 {
					dp[j][i] = true
				} else {
					dp[j][i] = dp[j+1][i-1]
				}
			}
		}
	}
	max := 0
	res := ""
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] {
				if j-i+1 > max {
					max = j - i + 1
					res = s[i : j+1]
				}
			}
		}
	}
	return res
}
