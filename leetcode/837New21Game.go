package leetcode

func new21Game(n int, k int, maxPts int) float64 {
	dp := make([]float64, n+1)
	dp[0] = 1.0
	start := float64(0)
	if k > 0 {
		start = float64(1)
	}
	for i := 1; i <= n; i++ {
		dp[i] = start / float64(maxPts)
		if i < k {
			start += dp[i]
		}
		if i-maxPts >= 0 && i-maxPts < k {
			start -= dp[i-maxPts]
		}
	}
	ans := float64(0)
	for i := k; i <= n; i++ {
		ans += dp[i]
	}
	return ans
}
