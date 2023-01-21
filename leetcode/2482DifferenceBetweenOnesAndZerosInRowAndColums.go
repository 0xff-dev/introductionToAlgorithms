package leetcode

func onesMinusZeros(grid [][]int) [][]int {
	r, c := len(grid), len(grid[0])
	ans := make([][]int, r)
	rows := make([]int, len(grid))
	cols := make([]int, len(grid[0]))
	for row := 0; row < r; row++ {
		ans[row] = make([]int, c)
		for col := 0; col < c; col++ {
			rows[row] += grid[row][col]
			cols[col] += grid[row][col]
		}
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			ans[i][j] = rows[i]*2 + cols[j]*2 - r - c
		}
	}
	return ans
}
