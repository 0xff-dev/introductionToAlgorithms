package leetcode

/*
func minInsertions(s string) int {
	length := len(s)
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
		for j := 0; j < length; j++ {
			dp[i][j] = -1
		}
		dp[i][i] = 0
	}
	for i := 1; i < length; i++ {
		for pre := i - 1; pre >= 0; pre-- {
			if s[pre] == s[i] {
				x := dp[pre+1][i-1]
				if pre == i-1 {
					x = 0
				}
				if dp[pre][i] == -1 || x < dp[pre][i] {
					dp[pre][i] = x
				}
			} else {
				if dp[pre][i] == -1 || dp[pre][i] > dp[pre][i-1]+1 {
					dp[pre][i] = dp[pre][i-1] + 1
				}
			}

			need := dp[pre][i]
			if pre >= 1 {
				x1 := need + pre
				x2 := dp[0][pre-1] + i - pre + 1
				need = x1
				if x2 < need {
					need = x2
				}
			}

			if dp[0][i] == -1 || need < dp[0][i] {
				dp[0][i] = need
			}
		}
	}
	return dp[0][length-1]
}
*/

func minInsertions(s string) int {
	length := len(s)
	one := []byte(s)
	for s, e := 0, length-1; s < e; s, e = s+1, e-1 {
		one[s], one[e] = one[e], one[s]
	}
	dp := make([][]int, length+1)
	for i := 0; i <= length; i++ {
		dp[i] = make([]int, length+1)
	}
	for row := 1; row <= length; row++ {
		for col := 1; col <= length; col++ {
			if one[col-1] == s[row-1] {
				dp[row][col] = dp[row-1][col-1] + 1
			} else {
				x := dp[row-1][col]
				if y := dp[row][col-1]; y > x {
					x = y
				}
				dp[row][col] = x
			}
		}
	}
	return length - dp[length][length]
}
