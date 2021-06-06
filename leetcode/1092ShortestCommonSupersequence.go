package leetcode

// 通过最长公共子序列，补充最少的字母
func shortestCommonSupersequence(str1 string, str2 string) string {
	l1, l2 := len(str1), len(str2)
	dp := make([][]int, l1+1)
	for idx := 0; idx <= l1; idx++ {
		dp[idx] = make([]int, l2+1)
	}

	for row := 1; row <= l1; row++ {
		for col := 1; col <= l2; col++ {
			if str1[row-1] == str2[col-1] {
				dp[row][col] = dp[row-1][col-1] + 1
				continue
			}
			dp[row][col] = dp[row][col-1]
			if dp[row-1][col] > dp[row][col] {
				dp[row][col] = dp[row-1][col]
			}
		}
	}
	x, y := l1, l2
	r := ""
	for x > 0 && y > 0 {
		if str1[x-1] == str2[y-1] {
			r = string(str1[x-1]) + r
			x, y = x-1, y-1
			continue
		}
		// 不想等
		left := dp[x][y-1]
		up := dp[x-1][y]
		if left < up {
			r = string(str1[x-1]) + r
			x--
			continue
		}
		r = string(str2[y-1]) + r
		y--
	}

	if x > 0 {
		r = str1[:x] + r
	} else if y > 0 {
		r = str2[:y] + r
	}
	return r
}
