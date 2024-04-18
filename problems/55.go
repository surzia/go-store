package problems

func canJump(nums []int) bool {
	var maxJump = 0

	for i := 0; i < len(nums); i++ {
		if maxJump < i {
			return false
		}

		if nums[i]+i > maxJump {
			maxJump = nums[i] + i
		}
	}
	return true
}
