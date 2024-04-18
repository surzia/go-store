package problems

func findPeakElement(nums []int) int {
	var left, right = 0, len(nums) - 1
	for left < right {
		var mid = left + (right-left)/2
		if (mid == len(nums)-1 || nums[mid] > nums[mid+1]) && (mid == 0 || nums[mid] > nums[mid-1]) {
			return mid
		}

		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
