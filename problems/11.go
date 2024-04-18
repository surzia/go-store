package problems

func maxArea(height []int) int {
	var n = len(height)
	var area = 0
	var left, right = 0, n - 1

	var max func(i, j int) int
	max = func(i, j int) int {
		if i > j {
			return i
		}

		return j
	}

	var min func(i, j int) int
	min = func(i, j int) int {
		if i > j {
			return j
		}

		return i
	}

	for left < right {
		area = max(min(height[left], height[right])*(right-left), area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return area
}
