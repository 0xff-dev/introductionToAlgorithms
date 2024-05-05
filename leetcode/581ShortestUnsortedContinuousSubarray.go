package leetcode

import "sort"

func findUnsortedSubarray(nums []int) int {
	// 2,6,4,8,10,9,15
	// 0,1,
	// 2,4,6,8,9,10,15
	dst := make([]int, len(nums))
	copy(dst, nums)
	sort.Ints(dst)
	l, r := 0, len(nums)-1
	for ; l <= r && nums[l] == dst[l]; l++ {
	}
	for ; r >= l && nums[r] == dst[r]; r-- {
	}
	return r - l + 1
}
