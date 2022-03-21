package sort

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func OptimizedBubbleSort(arr []int) {
	swap := true
	for i := 0; i < len(arr)-1; i++ {
		swap = false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
}

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
