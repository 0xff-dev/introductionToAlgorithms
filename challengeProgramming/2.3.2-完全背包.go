package challengeProgramming

import "fmt"

// 完全背包, 某个物品可以无限制的选择
func CompleteBackpack(weights, values []int, capacity int) int {
	lw, lv := len(weights), len(values)
	if lw != lv {
		fmt.Println("weights does't equal values")
		return 0
	}

	dp := make([]int, capacity+1)
	for goods := 0; goods < lw; goods++ {
		for _cap := capacity; _cap > 0; _cap-- {
			// 支持选择同一个物品多个, 0-n
			for cnt := 1; cnt*weights[goods] <= _cap; cnt++ {
				dp[_cap] = max(dp[_cap], dp[_cap-weights[goods]*cnt]+values[goods]*cnt)
			}
		}
	}
	return dp[capacity]
}
