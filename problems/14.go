package problems

func longestCommonPrefix(strs []string) string {
	var first = strs[0]
	if len(first) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return first
	}

	var index = 0
	var last = 0
	for {
		var flag = true
		for i := 1; i < len(strs); i++ {
			if index >= len(strs[i]) || index >= len(strs[0]) || strs[0][index] != strs[i][index] {
				flag = false
			}
			last = index
		}
		index++

		if !flag {
			break
		}
	}

	return first[:last]
}
