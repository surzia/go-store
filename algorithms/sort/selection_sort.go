package sort

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
