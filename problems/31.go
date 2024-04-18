package problems

func nextPermutation(nums []int) {
	var p = len(nums) - 1
	for p >= 1 && nums[p-1] >= nums[p] {
		p--
	}

	for i, j := p, len(nums)-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	if p == 0 {
		return
	}

	for i := p; i < len(nums); i++ {
		if nums[i] > nums[p-1] {
			nums[p-1], nums[i] = nums[i], nums[p-1]
			return
		}
	}
}
