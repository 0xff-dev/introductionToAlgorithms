package leetcode

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])
	firstRowHasZero := false
	firstColHasZero := false

	// 检查第一行是否有零
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstRowHasZero = true
			break
		}
	}

	// 检查第一列是否有零
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstColHasZero = true
			break
		}
	}

	// 使用第一行和第一列作为标记
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 根据标记将对应的行和列清零
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 如果第一行有零，全部清零
	if firstRowHasZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	// 如果第一列有零，全部清零
	if firstColHasZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
