package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rows, cols := len(obstacleGrid), len(obstacleGrid[0])

	dp := make([][]int, 2)
	for idx := 0; idx < 2; idx++ {
		dp[idx] = make([]int, cols)
	}
	if obstacleGrid[0][0] != 1 {
		dp[0][0] = 1
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if row == 0 && col == 0 {
				continue
			}
			if obstacleGrid[row][col] == 1 {
				dp[row&1][col] = 0
				continue
			}
			step := 0
			if row >= 1 {
				step += dp[1-row&1][col]
			}
			if col >= 1 {
				step += dp[row&1][col-1]
			}
			dp[row&1][col] = step
		}
	}

	return dp[1-rows&1][cols-1]
}
