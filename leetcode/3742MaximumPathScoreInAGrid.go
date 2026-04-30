package leetcode

import "math"

func maxPathScore(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])

	const INF = math.MinInt32

	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
			for c := range dp[i][j] {
				dp[i][j][c] = INF
			}
		}
	}

	dp[0][0][0] = 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for c := 0; c <= k; c++ {
				if dp[i][j][c] == INF {
					continue
				}

				if i+1 < m {
					val := grid[i+1][j]
					cost := 0
					if val != 0 {
						cost = 1
					}
					if c+cost <= k {
						if dp[i+1][j][c+cost] < dp[i][j][c]+val {
							dp[i+1][j][c+cost] = dp[i][j][c] + val
						}
					}
				}

				if j+1 < n {
					val := grid[i][j+1]
					cost := 0
					if val != 0 {
						cost = 1
					}
					if c+cost <= k {
						if dp[i][j+1][c+cost] < dp[i][j][c]+val {
							dp[i][j+1][c+cost] = dp[i][j][c] + val
						}
					}
				}
			}
		}
	}

	ans := INF
	for c := 0; c <= k; c++ {
		if dp[m-1][n-1][c] > ans {
			ans = dp[m-1][n-1][c]
		}
	}

	if ans < 0 {
		return -1
	}
	return ans
}
