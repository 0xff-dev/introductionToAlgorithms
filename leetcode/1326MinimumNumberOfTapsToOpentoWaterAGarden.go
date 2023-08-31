package leetcode

func minTaps(n int, ranges []int) int {
	const INF = 1000000000
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = INF
	}
	dp[0] = 0
	for i := 0; i <= n; i++ {
		left := 0
		if r := i - ranges[i]; r >= 0 {
			left = r
		}
		right := n
		if r := i + ranges[i]; r <= n {
			right = r
		}
		for j := left; j <= right; j++ {
			if dp[j]+1 < dp[right] {
				dp[right] = dp[j] + 1
			}
		}
	}
	if dp[n] == INF {
		return -1
	}
	return dp[n]
}
