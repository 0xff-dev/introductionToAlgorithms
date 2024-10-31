package leetcode

import "sort"

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	// Sort robots and factories by position
	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})

	// Flatten factory positions according to their capacities
	var factoryPositions []int
	for _, f := range factory {
		for i := 0; i < f[1]; i++ {
			factoryPositions = append(factoryPositions, f[0])
		}
	}

	robotCount := len(robot)
	factoryCount := len(factoryPositions)
	memo := make([][]int64, robotCount)
	for i := range memo {
		memo[i] = make([]int64, factoryCount)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	// Recursively calculate minimum total distance with memoization
	return calculateMinDistance(0, 0, robot, factoryPositions, memo)
}

func calculateMinDistance(robotIdx, factoryIdx int, robot, factoryPositions []int, memo [][]int64) int64 {
	// All robots assigned
	if robotIdx == len(robot) {
		return 0
	}
	// No factories left to assign
	if factoryIdx == len(factoryPositions) {
		return 1e12
	}
	// Check memo
	if memo[robotIdx][factoryIdx] != -1 {
		return memo[robotIdx][factoryIdx]
	}

	// Option 1: Assign current robot to current factory
	assign := int64(abs(robot[robotIdx]-factoryPositions[factoryIdx])) +
		calculateMinDistance(robotIdx+1, factoryIdx+1, robot, factoryPositions, memo)

	// Option 2: Skip current factory for the current robot
	skip := calculateMinDistance(robotIdx, factoryIdx+1, robot, factoryPositions, memo)

	memo[robotIdx][factoryIdx] = min(assign, skip)
	return memo[robotIdx][factoryIdx]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
