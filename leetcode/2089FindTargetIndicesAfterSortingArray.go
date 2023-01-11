package leetcode

import "sort"

func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	ans := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			ans = append(ans, i)
		}
	}
	return ans
}
