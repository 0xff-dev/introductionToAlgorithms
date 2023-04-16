package challengeprogrammingdatastructure

func Backpack(v, w []int, total int) int {
	dp := make([][]int, len(v)+1)
	for i := 0; i <= len(v); i++ {
		dp[i] = make([]int, total+1)
	}
	for i := 1; i <= len(v); i++ {
		for c := 1; c <= total; c++ {
			// more情况不装
			dp[i][c] = dp[i-1][c]
			if dp[i][c-1] > dp[i][c] {
				dp[i][c] = dp[i][c-1]
			}

			if c >= w[i-1] && dp[i-1][c-w[i-1]]+v[i-1] > dp[i][c] {
				dp[i][c] = dp[i-1][c-w[i-1]] + v[i-1]
			}
		}
	}
	return dp[len(v)][total]
}
