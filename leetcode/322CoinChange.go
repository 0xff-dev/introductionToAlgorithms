package leetcode

const minCost = -1

func specialMin(a, b int) int {
	if a == minCost || a > b {
		return b
	}
	return a
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	r := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		now := minCost
		for in := 0; in < len(coins); in++ {
			if i-coins[in] >= 0 && r[i-coins[in]] != minCost {
				now = specialMin(now, r[i-coins[in]]+1)
			}
		}
		r[i] = now
	}
	return r[amount]
}
