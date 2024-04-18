package problems

func removeDuplicates(nums []int) int {
	var left, right = 0, 0
	for right < len(nums) {
		for right < len(nums) && nums[left] == nums[right] {
			right++
		}

		if right >= len(nums) {
			break
		}
		left++
		nums[left], nums[right] = nums[right], nums[left]
		right++
	}
	return left + 1
}
