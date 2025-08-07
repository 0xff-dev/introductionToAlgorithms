package leetcode

import "math"

func maxCollectedFruits(fruits [][]int) int {
	n := len(fruits)
	ans := 0
	for i := 0; i < n; i++ {
		ans += fruits[i][i]
	}

	dp := func() int {
		prev := make([]int, n)
		curr := make([]int, n)
		for i := range prev {
			prev[i] = math.MinInt
		}
		prev[n-1] = fruits[0][n-1]
		for i := 1; i < n-1; i++ {
			for j := range curr {
				curr[j] = math.MinInt
			}
			for j := max(n-1-i, i+1); j < n; j++ {
				best := prev[j]
				if j-1 >= 0 {
					best = max(best, prev[j-1])
				}
				if j+1 < n {
					best = max(best, prev[j+1])
				}
				curr[j] = best + fruits[i][j]
			}
			prev, curr = curr, prev
		}
		return prev[n-1]
	}

	ans += dp()
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fruits[i][j], fruits[j][i] = fruits[j][i], fruits[i][j]
		}
	}
	ans += dp()
	return ans
}
