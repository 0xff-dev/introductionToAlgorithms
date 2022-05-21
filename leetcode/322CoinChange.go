package leetcode

const compareMin = -1

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for cost := 1; cost <= amount; cost++ {

		dp[cost] = compareMin
		for _, coin := range coins {
			if r := cost - coin; r >= 0 && dp[r] != compareMin {
				if dp[cost] == compareMin || dp[r]+1 < dp[cost] {
					dp[cost] = dp[r] + 1
				}
			}
		}
	}
	return dp[amount]
}
