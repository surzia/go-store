package problems

func maximalNetworkRank(n int, roads [][]int) int {
	if len(roads) == 0 {
		return 0
	}
	city := make(map[int]map[int]bool)
	for _, road := range roads {
		if _, ok := city[road[0]]; !ok {
			city[road[0]] = make(map[int]bool)
		}
		if _, ok := city[road[1]]; !ok {
			city[road[1]] = make(map[int]bool)
		}
		city[road[0]][road[1]] = true
		city[road[1]][road[0]] = true
	}
	a := make(map[int]int)
	for k, v := range city {
		a[k] = len(v)
	}

	max := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			var res int
			if _, ok := city[i][j]; ok {
				res = a[i] + a[j] - 1
			} else {
				res = a[i] + a[j]
			}

			if res > max {
				max = res
			}
		}
	}

	return max
}
