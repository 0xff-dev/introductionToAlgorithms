package offer

func Find(matrix [][]int, rows, cols, aim int) bool {
	if rows > 0 && cols > 0 {
		row, col := 0, cols-1
		for row < rows && col >= 0 {
			if matrix[row][col] == aim {
				return true
			}
			if matrix[row][col] < aim {
				row++
			} else {
				col--
			}
		}
	}
	return false
}
