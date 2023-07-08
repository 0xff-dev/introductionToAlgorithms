package leetcode

import "sort"

func putMarbles(weights []int, k int) int64 {
	l := len(weights)
	if l == k {
		return 0
	}
	/*
		dp := make([][]int, l)
		for i := 0; i < l; i++ {
			dp[i] = make([]int, k+1)
			dp[i][1] = weights[0] + weights[i]
		}
		// 10^14, 2^32
		for idx := 0; idx < l; idx++ {
			for group := 2; group <= k && group <= idx+1; group++ {
				for pre := idx; pre >= group-1; pre-- {
					cost := weights[pre] + weights[idx]
					preMax := dp[pre-1][group-1]
					if r := preMax + cost; r > dp[idx][group] {
						dp[idx][group] = r
					}
				}
			}
		}
		max := dp[l-1][k]
		for i := 0; i < l; i++ {
			for j := 0; j <= k; j++ {
				dp[i][j] = 0
			}
			dp[i][1] = weights[0] + weights[1]
		}
		for idx := 0; idx < l; idx++ {
			for group := 2; group <= k && group <= idx+1; group++ {
				for pre := idx; pre >= group-1; pre-- {
					cost := weights[pre] + weights[idx]
					preMin := dp[pre-1][group-1]
					if r := preMin + cost; dp[idx][group] == 0 || r < dp[idx][group] {
						dp[idx][group] = r
					}
				}
			}
		}
		return int64(max) - int64(dp[l-1][k])
	*/
	pairWeights := make([]int, l-1)
	for i := 0; i < l-1; i++ {
		pairWeights[i] += weights[i] + weights[i+1]
	}
	sort.Ints(pairWeights)
	ans := int64(0)
	for i := 0; i < k-1; i++ {
		ans += int64(pairWeights[l-2-1]) - int64(pairWeights[i])
	}
	return ans
}
