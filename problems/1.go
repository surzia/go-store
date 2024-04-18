package problems

func twoSum(nums []int, target int) []int {
	var hash = make(map[int]int)
	res := make([]int, 0, 2)
	for i := 0; i < len(nums); i++ {
		if k, ok := hash[nums[i]]; ok {
			res = append(res, k)
			res = append(res, i)
			break
		}
		hash[target-nums[i]] = i
	}
	return res
}
