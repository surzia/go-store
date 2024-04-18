package problems

func checkRecord(s string) bool {
	var aCount = 0
	var lCount = 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			aCount++
			if aCount > 1 {
				return false
			}
		}
		if s[i] == 'L' {
			lCount++
			if lCount > 2 {
				return false
			}
		} else {
			lCount = 0
		}
	}

	return true
}
