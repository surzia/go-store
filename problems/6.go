package problems

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	rows := make([][]string, numRows)
	direction := 1
	row := 0

	for _, char := range s {
		rows[row] = append(rows[row], string(char))
		if row == 0 {
			direction = 1
		} else if row == numRows-1 {
			direction = -1
		}
		row += direction
	}

	var result []string
	for _, r := range rows {
		result = append(result, r...)
	}

	return strings.Join(result, "")
}
