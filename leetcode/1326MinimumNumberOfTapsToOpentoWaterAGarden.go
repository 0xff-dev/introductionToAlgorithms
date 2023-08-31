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

func minTaps1(n int, ranges []int) int {
	cache := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		cache[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			cache[i][j] = -1
		}
	}
	// dfs + cache
	var dfs func(int, int) int
	dfs = func(left, right int) int {
		if cache[left][right] != -1 {
			return cache[left][right]
		}
		ans := -1
		for open := left; open <= right; open++ {
			if ranges[open] == 0 {
				continue
			}
			ll, rr := 0, 0
			l := open - ranges[open]
			r := open + ranges[open]
			if l > left {
				ll = dfs(left, l)
			}
			if r < right {
				rr = dfs(r, right)
			}
			if ll == -1 || rr == -1 {
				continue
			}
			if r := ll + rr + 1; ans == -1 || r < ans {
				ans = r
			}
		}
		cache[left][right] = ans
		return ans
	}
	dfs(0, n)
	return cache[0][n]
}
