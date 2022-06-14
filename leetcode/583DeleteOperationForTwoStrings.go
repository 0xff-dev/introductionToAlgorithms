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

func Solution(word1, word2 string) int {
	dp := make([][]int, 2)
	for idx := 0; idx < 2; idx++ {
		dp[idx] = make([]int, len(word2)+1)
	}
	for idx := 0; idx < len(word2)+1; idx++ {
		dp[0][idx] = idx
	}
	loop := 1
	for _, _byte := range []byte(word1) {
		dp[loop][0] = dp[1-loop][0] + 1
		for col := 1; col <= len(word2); col++ {
			leftTop := dp[1-loop][col-1]
			left := dp[loop][col-1]
			top := dp[1-loop][col]
			dp[loop][col] = leftTop
			if _byte != word2[col-1] {
				dp[loop][col] += 2
				if left+1 < dp[loop][col] {
					dp[loop][col] = left + 1
				}
				if top+1 < dp[loop][col] {
					dp[loop][col] = top + 1
				}
			}
		}

		loop = 1 - loop
	}
	return dp[1-loop][len(word2)]
}
