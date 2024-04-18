package problems

func permute(nums []int) [][]int {
	var res [][]int
	var p []int
	var flags []bool
	for i := 0; i < len(nums); i++ {
		flags = append(flags, false)
	}
	var df func(arr []int)
	df = func(arr []int) {
		if len(arr) == len(nums) {
			var cp []int
			for _, v := range arr {
				cp = append(cp, v)
			}
			res = append(res, cp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if flags[i] {
				continue
			}

			flags[i] = true
			arr = append(arr, nums[i])
			df(arr)
			arr = arr[:len(arr)-1]
			flags[i] = false
		}
	}
	df(p)

	return res
}
