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
