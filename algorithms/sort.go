// 在 algorithms 目录下创建 sort.go 文件

package algorithms

// 冒泡排序
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 插入排序
func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// 选择排序
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// 归并排序
func MergeSort(arr []int) {
	if len(arr) > 1 {
		mid := len(arr) / 2
		left := make([]int, mid)
		right := make([]int, len(arr)-mid)

		copy(left, arr[:mid])
		copy(right, arr[mid:])

		MergeSort(left)
		MergeSort(right)

		merge(arr, left, right)
	}
}

func merge(arr, left, right []int) {
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}

// 快速排序
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivot := arr[0]
	left, right := 1, len(arr)-1

	for left <= right {
		if arr[left] > pivot && arr[right] < pivot {
			arr[left], arr[right] = arr[right], arr[left]
		}
		if arr[left] <= pivot {
			left++
		}
		if arr[right] >= pivot {
			right--
		}
	}

	arr[0], arr[right] = arr[right], arr[0]

	QuickSort(arr[:right])
	QuickSort(arr[right+1:])
}

// 堆排序
func HeapSort(arr []int) {
	buildHeap(arr)

	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i)
	}
}

func buildHeap(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, i, n)
	}
}

func heapify(arr []int, idx, n int) {
	largest := idx
	left, right := 2*idx+1, 2*idx+2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != idx {
		arr[idx], arr[largest] = arr[largest], arr[idx]
		heapify(arr, largest, n)
	}
}

// 计数排序
func CountingSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	maxVal := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxVal {
			maxVal = arr[i]
		}
	}

	count := make([]int, maxVal+1)
	for _, v := range arr {
		count[v]++
	}

	i := 0
	for j, cnt := range count {
		for k := 0; k < cnt; k++ {
			arr[i] = j
			i++
		}
	}
}

// 桶排序
func BucketSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	maxVal, minVal := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxVal {
			maxVal = arr[i]
		}
		if arr[i] < minVal {
			minVal = arr[i]
		}
	}

	bucketSize := 5 // 可根据需要调整
	bucketCount := (maxVal-minVal)/bucketSize + 1

	buckets := make([][]int, bucketCount)
	for i := 0; i < bucketCount; i++ {
		buckets[i] = []int{}
	}

	for _, v := range arr {
		buckets[(v-minVal)/bucketSize] = append(buckets[(v-minVal)/bucketSize], v)
	}

	index := 0
	for _, bucket := range buckets {
		CountingSort(bucket)
		for _, v := range bucket {
			arr[index] = v
			index++
		}
	}
}

// 基数排序
func RadixSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	maxVal := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxVal {
			maxVal = arr[i]
		}
	}

	for exp := 1; maxVal/exp > 0; exp *= 10 {
		countSortByDigit(arr, exp)
	}
}

func countSortByDigit(arr []int, exp int) {
	n := len(arr)
	count := make([]int, 10)
	output := make([]int, n)

	for i := 0; i < n; i++ {
		count[(arr[i]/exp)%10]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		output[count[(arr[i]/exp)%10]-1] = arr[i]
		count[(arr[i]/exp)%10]--
	}

	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
}
