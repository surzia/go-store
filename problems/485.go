package problems

func findMaxConsecutiveOnes(nums []int) int {
	var dp = make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			dp[i] = 0
		} else {
			dp[i] = dp[i-1] + 1
		}
	}
	var max = 0
	for i := 0; i < len(dp); i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
