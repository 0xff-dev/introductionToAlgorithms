package leetcode

func luckyNumbers(matrix [][]int) []int {
	ans := make([]int, 0)
	m, n := len(matrix), len(matrix[0])
	minRows := make([]int, m)
	for row := 0; row < m; row++ {
		minRows[row] = 0
		for col := 1; col < n; col++ {
			if matrix[row][col] < matrix[row][minRows[row]] {
				minRows[row] = col
			}
		}
	}

	for col := 0; col < n; col++ {
		max := 0
		for row := 1; row < m; row++ {
			if matrix[row][col] > matrix[max][col] {
				max = row
			}
		}

		if minRows[max] == col {
			ans = append(ans, matrix[max][col])
		}
	}

	return ans
}
