package leetcode

func numMatchingSubseq(s string, words []string) int {
	r := 0
	for _, w := range words {
		if isSubseq(s, w) {
			r++
		}
	}
	return r
}

func isSubseq(t, s string) bool {
	it, is := 0, 0
	for it < len(t) && is < len(s) {
		if t[it] == s[is] {
			is++
		}
		it++
	}
	return is == len(s)
}
func canMatch(t, s string) bool {
	ls, lt := len(s), len(t)
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, ls+1)
	}

	loop := 1
	for row := 0; row < lt; row++ {
		for col := 1; col <= ls; col++ {
			if t[row] == s[col-1] {
				dp[loop][col] = dp[1-loop][col-1] + 1
				continue
			}
			m := dp[1-loop][col]
			if dp[loop][col-1] > m {
				m = dp[loop][col-1]
			}
			dp[loop][col] = m
		}
		loop = 1 - loop
	}
	return dp[1-loop][ls] == lt
}
