package leetcode

import "sort"

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	pre := nums[0]
	for idx := 1; idx < len(nums); idx++ {
		if nums[idx] == pre {
			return true
		}
		pre = nums[idx]
	}
	return false
}
