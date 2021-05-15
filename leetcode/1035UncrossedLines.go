package leetcode

func maxUncrossedLines(A []int, B []int) int {
	l1, l2 := len(A), len(B)

	dp := [2][]int{}
	for idx := 0; idx < 2; idx++ {
		dp[idx] = make([]int, l2+1)
	}

	loop := 1
	for idx := 0; idx < l1; idx++ {
		for inner := 1; inner <= l2; inner++ {
			if A[idx] == B[inner-1] {
				dp[loop][inner] = dp[1-loop][inner-1] + 1
				continue
			}

			dp[loop][inner] = dp[1-loop][inner]
			if dp[loop][inner-1] > dp[loop][inner] {
				dp[loop][inner] = dp[loop][inner-1]
			}
		}
		loop = 1 - loop
	}
	return dp[1-loop][l2]
}
