package leetcode

const mod920 = 1000000007

func numMusicPlaylists(n int, goal int, k int) int {
	dp := make([][]int, goal+1)
	for i := 0; i <= goal; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1
	for i := 1; i <= goal; i++ {
		end := i
		if n < end {
			end = n
		}
		for j := 1; j <= end; j++ {
			dp[i][j] = dp[i-1][j-1] * (n - j + 1) % mod920
			if j > k {
				dp[i][j] = (dp[i][j] + dp[i-1][j]*(j-k)) % mod920
			}
		}
	}
	return dp[goal][n]
}
