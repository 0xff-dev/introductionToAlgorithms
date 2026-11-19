package leetcode

import "sort"

func minimizeSum(nums []int) int {
	// 绝对值最小的 改成相同肯定就是0
	// 绝对值最大的
	// change [0,1], [n-1, n-2], [0, n-1]
	n := len(nums)
	sort.Ints(nums)
	a := nums[n-1] - nums[2]
	b := nums[n-2] - nums[1]
	c := nums[n-3] - nums[0]
	return min(a, b, c)
}
