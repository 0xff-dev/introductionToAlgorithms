package leetcode

func uniquePaths(m int, n int) int {
	r := make([][]int, m)
	for idx := 0; idx < m; idx++ {
		r[idx] = make([]int, n)
		r[idx][0] = 1
		if idx == 0 {
			for i := 0; i < n; i++ {
				r[idx][i] = 1
			}
		}
	}

	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			r[row][col] = r[row][col-1] + r[row-1][col]
		}
	}

	return r[m-1][n-1]
}
