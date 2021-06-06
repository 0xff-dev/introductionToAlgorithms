package leetcode

func longestPalindromeSubseq(s string) int {
	ls := len(s)
	reverse := make([]byte, ls)
	for idx := len(s) - 1; idx >= 0; idx-- {
		reverse[ls-1-idx] = s[idx]
	}

	dp := make([][]int, 2)
	for idx := 0; idx < 2; idx++ {
		dp[idx] = make([]int, ls+1)
	}

	loop := 1
	for row := 0; row < ls; row++ {
		for col := 1; col <= ls; col++ {
			if reverse[row] == s[col-1] {
				dp[loop][col] = dp[1-loop][col-1] + 1
				continue
			}

			_max := dp[1-loop][col]
			if dp[loop][col-1] > _max {
				_max = dp[loop][col-1]
			}
			dp[loop][col] = _max
		}

		loop = 1 - loop
	}

	return dp[1-loop][ls]
}
