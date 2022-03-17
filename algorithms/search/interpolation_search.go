package search

// InterpolationSearch is an improvement over binary search tailored for
// uniformly distributed data. Binary search halves the search space on
// each step regardless of the data distribution, thus it's time complexity
// is always O(log(n)). On the other hand, interpolation search time complexity
// varies depending on the data distribution. It is faster than binary search
// for uniformly distributed data with the time complexity of O(log(log(n))).
// However, in the worst-case scenario, it can perform as poor as O(n).
func InterpolationSearch(array []int, key int) int {
	highEnd := len(array) - 1
	lowEnd := 0

	for {
		if key < array[lowEnd] || key > array[highEnd] || lowEnd > highEnd {
			break
		}

		probe := lowEnd + (highEnd-lowEnd)*(key-array[lowEnd])/(array[highEnd]-array[lowEnd])
		if highEnd == lowEnd {
			if array[lowEnd] == key {
				return lowEnd
			}
			return -1
		}

		if array[probe] == key {
			return probe
		} else if array[probe] < key {
			lowEnd = probe + 1
		} else {
			highEnd = probe - 1
		}
	}

	return -1
}
