package problems

func kthDistinct(arr []string, k int) string {
	var hash = make(map[string]int)
	for i := 0; i < len(arr); i++ {
		if v, ok := hash[arr[i]]; ok {
			hash[arr[i]] = v + 1
		} else {
			hash[arr[i]] = 1
		}
	}

	var res = ""
	for i := 0; i < len(arr); i++ {
		if hash[arr[i]] == 1 {
			res = arr[i]
			k--
			if k == 0 {
				return res
			}
			res = ""
		}
	}
	return res
}
