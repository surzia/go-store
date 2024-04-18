package problems

func subarraySum(nums []int, k int) int {
	var prefix = make(map[int]int)
	var res = 0
	var sum = 0
	prefix[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if v, ok := prefix[sum-k]; ok {
			res += v
		}
		if v, ok := prefix[sum]; ok {
			prefix[sum] = v + 1
		} else {
			prefix[sum] = 1
		}
	}
	return res
}
