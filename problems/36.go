package problems

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		var p = []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
		for j := 0; j < 9; j++ {
			if board[i][j] == byte('.') {
				continue
			}
			p[int(board[i][j]-'1')]--
			if p[int(board[i][j]-'1')] < 0 {
				return false
			}
		}
	}

	for i := 0; i < 9; i++ {
		var p = []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
		for j := 0; j < 9; j++ {
			if board[j][i] == byte('.') {
				continue
			}
			p[int(board[j][i]-'1')]--
			if p[int(board[j][i]-'1')] < 0 {
				return false
			}
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			var p = []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
			for m := i * 3; m < (i+1)*3; m++ {
				for n := j * 3; n < (j+1)*3; n++ {
					if board[m][n] == byte('.') {
						continue
					}
					p[int(board[m][n]-'1')]--
					if p[int(board[m][n]-'1')] < 0 {
						return false
					}
				}
			}
		}
	}
	return true
}
