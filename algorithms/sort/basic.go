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

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left := MergeSort(arr[:len(arr)/2])
	right := MergeSort(arr[len(arr)/2:])

	var final []int
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			final = append(final, left[i])
			i++
		} else {
			final = append(final, right[j])
			j++
		}
	}
	for ; i < len(left); i++ {
		final = append(final, left[i])
	}
	for ; j < len(right); j++ {
		final = append(final, right[j])
	}
	return final
}

func InsertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {

			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minElementIndex := i

		for j := i + 1; j < len(arr); j++ {
			if arr[minElementIndex] > arr[j] {
				minElementIndex = j
			}
		}

		if minElementIndex != i {
			arr[i], arr[minElementIndex] = arr[minElementIndex], arr[i]
		}
	}
}

func ShellSort(arr []int) {
	for i := len(arr) / 2; i > 0; i /= 2 {
		for j := i; j < len(arr); j++ {
			key := arr[j]
			k := j

			for k >= i && arr[k-i] > key {
				arr[k] = arr[k-i]
				k -= i
			}

			arr[k] = key
		}
	}
}
