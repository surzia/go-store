package problems

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var step = make([]int, len(nums))
	step[0] = 1
	var localMax = nums[0]
	var globalMax = nums[0]
	for i := 1; i < len(nums); i++ {
		if i <= localMax {
			step[i] = step[i-1]
		} else {
			step[i] = step[i-1] + 1
			localMax = globalMax
		}
		if globalMax < nums[i]+i {
			globalMax = nums[i] + i
		}
	}
	return step[len(nums)-1]
}
