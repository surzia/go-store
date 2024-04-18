package problems

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	left, right := 0, 0
	max := 0
	hash := make(map[string]bool)
	for right < len(s) {
		if _, ok := hash[string(s[right])]; !ok {
			hash[string(s[right])] = true
			if max < right-left+1 {
				max = right - left + 1
			}
			right++
		} else {
			delete(hash, string(s[left]))
			left++
		}
	}
	return max
}
