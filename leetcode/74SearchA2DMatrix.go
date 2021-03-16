package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	x, y := 0, len(matrix[0])-1
	for x > -1 && x < len(matrix) && y > -1 && y < len(matrix[0]) {
		num := matrix[x][y]
		if num == target {
			return true
		}
		if num < target {
			x++
			continue
		}
		if num > target {
			y--
		}
	}
	return false
}
