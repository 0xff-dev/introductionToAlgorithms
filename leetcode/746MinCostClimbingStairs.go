package leetcode

func minCostClimbingStairs(cost []int) int {
	oneStep, twoStep := cost[0], cost[1]
	for idx := 2; idx < len(cost); idx++ {
		_min := oneStep + cost[idx]
		if _min > twoStep+cost[idx] {
			_min = twoStep + cost[idx]
		}
		oneStep = twoStep
		twoStep = _min
	}
	if oneStep > twoStep {
		return twoStep
	}
	return oneStep
}
