package leetcode

import "slices"

func minMoves3736(nums []int) int {
	var ret int
	maxElem := slices.Max(nums)
	for i := range nums {
		ret += (maxElem - nums[i])
	}
	return ret
}
