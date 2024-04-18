package problems

import (
	"math"
	"sort"
)

func minMoves2(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	sort.Ints(nums)
	var move = math.MaxInt
	for i := 0; i < len(nums)-1; i++ {
		var x = nums[i]
		var y = 0
		for j := 0; j <= i; j++ {
			y += x - nums[j]
		}
		for j := i + 1; j < len(nums); j++ {
			y += nums[j] - x
		}
		if move > y {
			move = y
		}
	}
	return move
}
