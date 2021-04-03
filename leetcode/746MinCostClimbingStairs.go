package leetcode

func minCostClimbingStairs(cost []int) int {
	oneStep, twoStep := cost[0], cost[1]
	for idx := 2; idx < len(cost); idx++ {
		_min := oneStep
		if _min > twoStep {
			_min = twoStep
		}
		oneStep = twoStep
		twoStep = _min + cost[idx]
	}
	if oneStep > twoStep {
		return twoStep
	}
	return oneStep
}
