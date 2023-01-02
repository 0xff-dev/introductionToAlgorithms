package leetcode

func numDecodings(s string) int {
	l := len(s)
	dp := make([]int, l)
	if s[0] == '0' {
		return 0
	}
	dp[0] = 1
	for idx := 1; idx < l; idx++ {
		if s[idx] == '0' {
			if s[idx-1] != '1' && s[idx-1] != '2' {
				return 0
			}
			i := idx - 2
			if i < 0 {
				i = 0
			}
			dp[idx] = dp[i]
			continue
		}

		dp[idx] = dp[idx-1]
		if s[idx-1] == '0' {
			continue
		}
		if (s[idx-1] == '1') || s[idx-1] == '2' && s[idx] >= '1' && s[idx] <= '6' {
			i := idx - 2
			if i < 0 {
				i = 0
			}
			dp[idx] += dp[i]
		}
	}

	return dp[l-1]
}
