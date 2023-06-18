package leetcode

const mod2328 = 1000000007

func countPaths(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	dp := make([][]int, rows)
	/*
		for r := 0; r < rows; r++ {
			dp[r] = make([]int, cols)
			for i := 0; i < cols; i++ {
				dp[r][i] = 1
			}
		}
		// 左上角到右下角
		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				x, y := row-1, col
				if !(x < 0 || x >= rows || y < 0 || y >= cols || grid[row][col] <= dp[x][y]) {
					dp[row][col] = (dp[row][col]%mod2328 + dp[x][y]%mod2328) % mod2328
				}
				x, y = row, col-1
				if !(x < 0 || x >= rows || y < 0 || y >= cols || grid[row][col] <= dp[x][y]) {
					dp[row][col] = (dp[row][col]%mod2328 + dp[x][y]%mod2328) % mod2328
				}
			}
		}
		for row := rows - 1; row >= 0; row-- {
			for col := cols - 1; col >= 0; col-- {
				x, y := row-1, col
				if !(x < 0 || x >= rows || y < 0 || y >= cols || grid[row][col] <= dp[x][y]) {
					dp[row][col] = (dp[row][col]%mod2328 + dp[x][y]%mod2328) % mod2328
				}
				x, y = row, col-1
				if !(x < 0 || x >= rows || y < 0 || y >= cols || grid[row][col] <= dp[x][y]) {
					dp[row][col] = (dp[row][col]%mod2328 + dp[x][y]%mod2328) % mod2328
				}

			}
		}

		ans := 0
		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				ans = (ans + dp[row][col]) % mod2328
			}
		}
		return ans
	*/
	// try dfs
	var dirs = [][]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
	}
	for row := 0; row < rows; row++ {
		dp[row] = make([]int, cols)
	}
	var dfs func(int, int) int
	dfs = func(x, y int) int {
		if dp[x][y] != 0 {
			return dp[x][y]
		}
		// self
		count := 1
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx < 0 || nx >= rows || ny < 0 || ny >= cols || grid[nx][ny] <= grid[x][y] {
				continue
			}
			count += dfs(nx, ny)
		}
		dp[x][y] += count % mod2328
		return count
	}
	ans := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			ans += dfs(row, col)
			ans %= mod2328
		}
	}
	return ans
}
