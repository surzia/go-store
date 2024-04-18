package problems

func search(nums []int, target int) int {
	var df func() int
	df = func() int {
		var left, right = 0, len(nums) - 1
		for left <= right {
			var mid = left + (right-left)/2
			if mid == left {
				if nums[right] > nums[left] {
					return left
				} else {
					return right
				}
			}

			if nums[mid] < nums[mid-1] {
				return mid
			}

			if nums[mid] > nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return 0
	}
	var mod = df()

	var bs func() int
	bs = func() int {
		var left, right = 0, len(nums) - 1
		for left <= right {
			var mid = left + (right-left)/2
			var r = (mid + mod) % len(nums)
			if nums[r] == target {
				return r
			} else if nums[r] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return -1
	}
	return bs()
}
