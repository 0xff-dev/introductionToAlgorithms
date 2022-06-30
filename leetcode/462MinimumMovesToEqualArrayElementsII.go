package leetcode

import "sort"

/*
4, 9, 10, 12
4-4 9-4 + 10-4+12-4 --> 9-4 + 12-10
*/
func minMoves2(nums []int) int {
	length := len(nums)
	sort.Ints(nums)
	steps := 0
	for idx := 0; idx < length/2; idx++ {
		steps += nums[length-idx-1] - nums[idx]
	}
	return steps
}
