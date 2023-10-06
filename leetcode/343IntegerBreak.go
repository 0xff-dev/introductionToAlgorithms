package leetcode

func integerBreak(n int) int {
	dp := make([]int, 58)
	dp[1], dp[2], dp[3], dp[4] = 1, 1, 2, 4
	for target := 5; target <= n; target++ {
		// 1,4, 2,3, 3,2, 4,1
		for walker := 1; walker <= target/2; walker++ {
			max := -1
			if x1 := walker * (target - walker); x1 > max {
				max = x1
			}
			if x2 := walker * dp[target-walker]; x2 > max {
				max = x2
			}
			if x3 := dp[walker] * (target - walker); x3 > max {
				max = x3
			}
			if x4 := dp[walker] * dp[target-walker]; x4 > max {
				max = x4
			}
			if dp[target] < max {
				dp[target] = max
			}
		}
	}
	return dp[n]
}

func max343(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func integerBreak343(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		end := i / 2
		if i&1 == 1 {
			end++
		}
		for start := i - 1; start >= end; start-- {
			a := max343(start, dp[start])
			b := max343(i-start, dp[i-start])
			if a*b > dp[i] {
				dp[i] = a * b
			}
		}
	}
	return dp[n]
}
