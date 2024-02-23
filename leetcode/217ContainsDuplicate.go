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

func containsDuplicate1(nums []int) bool {
	c := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := c[n]; ok {
			return true
		}
		c[n] = struct{}{}
	}
	return false
}
