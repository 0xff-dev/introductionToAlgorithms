package leetcode

func isSubsequence(s string, t string) bool {
	l1, l2 := len(s), len(t)
	if l1 == 0 {
		return true
	}

	if l2 == 0 {
		return false
	}

	dp := [2][]int{}
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, l2+1)
	}

	loop := 1
	for row := 0; row < l1; row++ {
		for col := 1; col <= l2; col++ {
			if s[row] == t[col-1] {
				// 有一个相等
				dp[loop][col] = dp[1-loop][col-1] + 1
				continue
			}
			dp[loop][col] = dp[1-loop][col]
			if dp[loop][col-1] > dp[loop][col] {
				dp[loop][col] = dp[loop][col-1]
			}
		}

		loop = 1 - loop
	}

	return dp[1-loop][l2] == l1
}
