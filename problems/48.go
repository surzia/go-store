package problems

func rotate(matrix [][]int) {
	var rows = len(matrix)
	var cols = len(matrix[0])
	for i := 0; i < cols; i++ {
		for j := 0; j < rows/2; j++ {
			matrix[j][i], matrix[rows-j-1][i] = matrix[rows-j-1][i], matrix[j][i]
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
