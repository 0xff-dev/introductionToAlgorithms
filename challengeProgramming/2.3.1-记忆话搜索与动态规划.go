package challengeProgramming

/*
01背包问题
*/

func Backpack(weights, values []int, capacity int) int {
	if len(weights) != len(values) {
		return -1
	}
	dp := make([]int, capacity+1)
	for index, w := range weights {
		for v := capacity; v >= 0; v-- {
			if v >= w {
				dp[v] = max(dp[v], dp[v-w]+values[index])
			}
		}
	}
	return dp[capacity]
}
