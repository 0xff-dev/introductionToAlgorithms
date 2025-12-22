package leetcode

func minDeletionSize960(A []string) int {
	if len(A) == 0 {
		return 0
	}

	numCols := len(A[0])
	// dp[i] 表示以第 i 列结尾的最长保序子序列的长度
	dp := make([]int, numCols)

	for i := range dp {
		dp[i] = 1
	}

	// 从右向左计算 DP（与原 Java 代码逻辑一致）
	for i := numCols - 2; i >= 0; i-- {
		for j := i + 1; j < numCols; j++ {
			// 检查是否所有字符串在第 i 列的字符都 <= 第 j 列的字符
			isSorted := true
			for _, row := range A {
				if row[i] > row[j] {
					isSorted = false
					break
				}
			}

			if isSorted {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	maxKept := 0
	for _, val := range dp {
		if val > maxKept {
			maxKept = val
		}
	}

	return numCols - maxKept
}
