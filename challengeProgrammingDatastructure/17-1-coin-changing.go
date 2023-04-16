package challengeprogrammingdatastructure

import "sort"

// 有n种硬币, 硬币无限使用
func Coins(coins []int, target int) int {
	sort.Ints(coins)

	dp := make([]int, target+1)
	dp[1] = 1
	for coin := 2; coin <= target; coin++ {
		for _, c := range coins {
			if c > coin {
				break
			}
			if r := dp[coin-c] + 1; dp[coin] == 0 || r < dp[coin] {
				dp[coin] = r
			}
		}
	}
	return dp[target]
}
