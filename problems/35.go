package problems

func searchInsert(nums []int, target int) int {
	var left, right = 0, len(nums) - 1
	for left <= right {
		var mid = left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right + 1
}
