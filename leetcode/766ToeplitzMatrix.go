package leetcode

func isToeplitzMatrix(matrix [][]int) bool {
	m, n := len(matrix), len(matrix[0])
	for col := 0; col < n-1; col++ {
		for x, y := 1, col+1; x < m && y < n; x, y = x+1, y+1 {
			if matrix[x][y] != matrix[0][col] {
				return false
			}
		}
	}
	for row := 1; row < m-1; row++ {
		for x, y := row+1, 1; x < m && y < n; x, y = x+1, y+1 {
			if matrix[x][y] != matrix[row][0] {
				return false
			}
		}
	}
	return true
}
