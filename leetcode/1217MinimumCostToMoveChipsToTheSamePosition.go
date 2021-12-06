package leetcode

func minCostToMoveChips(position []int) int {
	minCost := -1
	for cmpIdx := 0; cmpIdx < len(position); cmpIdx++ {
		cost := 0
		for in := 0; in < len(position); in++ {
			absLen := position[in] - position[cmpIdx]
			if absLen < 0 {
				absLen = -absLen
			}
			if absLen&1 == 1 {
				cost++
			}
		}
		if minCost == -1 || minCost > cost {
			minCost = cost
		}
	}
	return minCost
}
