package sort

func QuickSort(arr []int, begin, end int) {
	if begin >= end {
		return
	}

	left, right := begin, end
	pivot := arr[left]

	for left < right {

		for left < right && arr[right] >= pivot {
			right--
		}
		if left < right {
			arr[left] = arr[right]
		}

		for left < right && arr[left] <= pivot {
			left++
		}
		if left < right {
			arr[right] = arr[left]
		}

		if left >= right {
			arr[left] = pivot
		}
	}

	QuickSort(arr, begin, right-1)
	QuickSort(arr, right+1, end)
}
