package problems

func magicalString(n int) int {
	if n == 1 {
		return 1
	}
	var str = []int{1}
	var step = 1
	var count = 0
	var p = 2
	var start = 2
	var last = 2
	for step < n {
		for i := 0; i < p; i++ {
			str = append(str, start)
		}
		p = str[last]
		last++
		step = len(str)
		if start == 1 {
			start = 2
		} else {
			start = 1
		}
	}
	for i := 0; i < n; i++ {
		if str[i] == 1 {
			count++
		}
	}
	return count
}
