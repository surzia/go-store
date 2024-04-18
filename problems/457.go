package problems

func circularArrayLoop(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		var hash = make(map[int]int)
		var circle = false
		var curIndex = i
		for {
			var step = nums[curIndex]
			var direction = 1
			if step < 0 {
				step = -step
				direction = -1
			}
			step = step % len(nums)
			var nextIndex = (curIndex + step*direction + len(nums)) % len(nums)
			if nextIndex == curIndex {
				break
			}
			if nums[curIndex] > 0 {
				hash[curIndex] = 1
			} else {
				hash[curIndex] = -1
			}
			_, exist := hash[nextIndex]
			if exist {
				if len(hash) > 1 {
					var flag = true
					for _, v := range hash {
						if v != hash[curIndex] {
							flag = false
						}
					}
					circle = flag
				}
				break
			}
			if nums[nextIndex] > 0 {
				hash[nextIndex] = 1
			} else {
				hash[nextIndex] = -1
			}
			curIndex = nextIndex
		}
		if circle {
			return true
		}
	}

	return false
}
