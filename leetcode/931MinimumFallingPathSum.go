package leetcode

const minSteps1 = 10001

func minFallingPathSum1(matrix [][]int) int {
	length := len(matrix)
	if length == 0 {
		return 0
	}

	dp := [2][]int{}
	for i := 0; i < 2; i++ {
		if i == 0 {
			dp[i] = matrix[0]
			continue
		}
		dp[i] = make([]int, length)
		for j := 0; j < length; j++ {
			dp[i][j] = minSteps1
		}
	}

	loop := 1
	for row := 1; row < length; row++ {
		// 上一个计算结果
		for col := 0; col < length; col++ {
			if col >= 1 {
				tmp := dp[1-loop][col] + matrix[row][col-1]
				if tmp < dp[loop][col-1] {
					dp[loop][col-1] = tmp
				}
			}

			tmp := dp[1-loop][col] + matrix[row][col]
			if tmp < dp[loop][col] {
				dp[loop][col] = tmp
			}

			if col+1 < length {
				tmp := dp[1-loop][col] + matrix[row][col+1]
				if tmp < dp[loop][col+1] {
					dp[loop][col+1] = tmp
				}
			}
			dp[1-loop][col] = minSteps1
		}
		loop = 1 - loop
	}
	r := dp[1-loop][0]

	for i := 1; i < length; i++ {
		if dp[1-loop][i] < r {
			r = dp[1-loop][i]
		}
	}
	return r
}
