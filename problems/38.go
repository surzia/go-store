package problems

import "strconv"

func countAndSay(n int) string {
	var df func(x int) []int
	df = func(x int) []int {
		if x == 1 {
			return []int{1}
		}

		var pre = df(x - 1)
		var cur []int
		var fast, slow = 0, 0
		for {
			for fast < len(pre) && pre[fast] == pre[slow] {
				fast++
			}

			if fast >= len(pre) {
				break
			}
			cur = append(cur, fast-slow)
			cur = append(cur, pre[slow])
			slow = fast
		}
		cur = append(cur, fast-slow)
		cur = append(cur, pre[slow])
		return cur
	}
	var res = df(n)
	var str string
	for _, r := range res {
		str += strconv.Itoa(r)
	}
	return str
}
