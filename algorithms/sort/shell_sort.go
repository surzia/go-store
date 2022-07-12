package sort

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
