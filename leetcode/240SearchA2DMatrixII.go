package leetcode

func searchMatrixII(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	x, y := 0, cols-1
	for x < rows && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] < target {
			x++
			continue
		}
		y--
	}
	return false
}
