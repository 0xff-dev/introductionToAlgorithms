package leetcode

import "sort"

// 感觉像是dp的问题
func minimizeMax(nums []int, p int) int {
	// 先通过dfs实现
	sort.Ints(nums)
	l := len(nums)
	left, right := 0, nums[l-1]-nums[0]
	var valid func(int) int
	valid = func(threshold int) int {
		index, count := 0, 0
		for ; index < l-1; index++ {
			if nums[index+1]-nums[index] <= threshold {
				count++
				index++
			}
		}
		return count
	}
	for left < right {
		mid := left + (right-left)/2
		if valid(mid) >= p {
			right = mid
			continue
		}
		left = mid + 1
	}
	return left
}
