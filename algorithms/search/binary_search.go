package search

// IterativeBinarySearch is iterative implementation of binary search.
// it searches key in array range low to high, if there is key in array,
// then return index of key. Otherwise, returns -1
func IterativeBinarySearch(array []int, key, low, high int) int {
	for {
		if low > high {
			break
		}

		mid := low + (high-low)/2
		if array[mid] == key {
			return mid
		} else if array[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// RecursiveBinarySearch is recursive implementation of binary search.
// it searches key in array range low to high, if there is key in array,
// then return index of key. Otherwise, returns -1
func RecursiveBinarySearch(array []int, key, low, high int) int {
	mid := low + (high-low)/2

	if low > high {
		return -1
	}

	if key == array[mid] {
		return mid
	} else if array[mid] < key {
		return RecursiveBinarySearch(array, key, mid+1, high)
	} else {
		return RecursiveBinarySearch(array, key, low, mid-1)
	}
}
