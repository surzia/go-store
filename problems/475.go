package problems

import "sort"

func findRadius(houses []int, heaters []int) int {
	var r = 0
	sort.Ints(heaters)
	var df func(x int, arr []int, left, right int) []int
	df = func(x int, arr []int, left, right int) []int {
		if left == right {
			return []int{left, left}
		}
		if left+1 == right {
			return []int{left, right}
		}

		var mid = (left + right) / 2
		if arr[mid] == x {
			return []int{mid, mid}
		} else if arr[mid] > x {
			right = mid
		} else {
			left = mid
		}
		return df(x, arr, left, right)
	}
	var abs = func(x, y int) int {
		if x > y {
			return x - y
		} else {
			return y - x
		}
	}
	var cover = make(map[int]int)
	for i := 0; i < len(houses); i++ {
		var index = df(houses[i], heaters, 0, len(heaters)-1)
		if abs(heaters[index[0]], houses[i]) < abs(heaters[index[1]], houses[i]) {
			cover[houses[i]] = abs(heaters[index[0]], houses[i])
		} else {
			cover[houses[i]] = abs(heaters[index[1]], houses[i])
		}
	}
	for _, v := range cover {
		if r < v {
			r = v
		}
	}
	return r
}
