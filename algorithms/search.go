package algorithms

// 线性搜索
func LinearSearch(arr []int, target int) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1
}

// 二分搜索
func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// 广度优先搜索 (BFS) - 这里仅作为示例，实际情况可能更复杂
func BFS(graph map[int][]int, start int, target int) bool {
	queue := []int{start}
	visited := make(map[int]bool)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == target {
			return true
		}

		if !visited[node] {
			visited[node] = true
			queue = append(queue, graph[node]...)
		}
	}

	return false
}

// 深度优先搜索 (DFS)
func DFS(graph map[int][]int, start int, target int) bool {
	visited := make(map[int]bool)
	return dfsRecursive(graph, start, target, visited)
}

func dfsRecursive(graph map[int][]int, node, target int, visited map[int]bool) bool {
	if node == target {
		return true
	}
	visited[node] = true
	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			if dfsRecursive(graph, neighbor, target, visited) {
				return true
			}
		}
	}
	return false
}
