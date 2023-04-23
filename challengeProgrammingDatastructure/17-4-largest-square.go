package challengeprogrammingdatastructure

func LargestSquare(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	dp := make([][]int, rows)
	for r := 0; r < rows; r++ {
		dp[r] = make([]int, cols)
	}
	for i := 0; i < cols; i++ {
		dp[0][i] = 1 - grid[0][i]
	}
	ans := 1
	for row := 1; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 1 {
				continue
			}
			if col == 0 {
				dp[row][col] = 1
				continue
			}
			x := dp[row-1][col]
			if y := dp[row][col-1]; x > y {
				x = y
			}
			if z := dp[row-1][col-1]; x > z {
				x = z
			}
			dp[row][col] = x + 1
			if dp[row][col] > ans {
				ans = dp[row][col]
			}
		}
	}
	return ans * ans
}
