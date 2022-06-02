package leetcode

func transpose(matrix [][]int) [][]int {
	rows, cols := len(matrix), len(matrix[0])
	if rows == cols {
		for idx := 0; idx < rows; idx++ {
			for j := idx + 1; j < cols; j++ {
				matrix[idx][j], matrix[j][idx] = matrix[j][idx], matrix[idx][j]
			}
		}
		return matrix
	}
	alloc := make([][]int, cols)
	for idx := 0; idx < cols; idx++ {
		alloc[idx] = make([]int, rows)
		for row := 0; row < rows; row++ {
			alloc[idx][row] = matrix[row][idx]
		}
	}
	return alloc
}
