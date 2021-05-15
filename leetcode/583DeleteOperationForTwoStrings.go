package leetcode

func minDistance1(word1 string, word2 string) int {
	word1Len, word2Len := len(word1), len(word2)
	if word1Len == 0 {
		return word2Len
	}
	if word2Len == 0 {
		return word1Len
	}

	dp := [2][]int{}

	for idx := 0; idx < 2; idx++ {
		dp[idx] = make([]int, word2Len+1)
		if idx == 0 {
			for col := 0; col < word2Len+1; col++ {
				dp[idx][col] = col
			}
			continue
		}
		dp[idx][0] = idx
	}

	loop := 1
	for row := 0; row < word1Len; row++ {
		for col := 1; col <= word2Len; col++ {
			if word1[row] == word2[col-1] {
				dp[loop][col] = dp[1-loop][col-1]
				continue
			}
			dp[loop][col] = dp[1-loop][col]
			if dp[loop][col-1] < dp[loop][col] {
				dp[loop][col] = dp[loop][col-1]
			}
			dp[loop][col]++
		}
		// dp 清零的动作,避免出现脏数据
		for col := 0; col < word2Len+1; col++ {
			dp[1-loop][col] = 0
		}

		dp[1-loop][0] = dp[loop][0] + 1
		loop = 1 - loop
	}
	return dp[1-loop][word2Len]
}
