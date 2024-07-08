package leetcode

import "sort"

func thirdMax(nums []int) int {
	distinct := make(map[int]struct{})
	index := 0
	for _, n := range nums {
		if _, ok := distinct[n]; !ok {
			distinct[n] = struct{}{}
			nums[index] = n
			index++
		}
	}
	nums = nums[:index]
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	if len(nums) < 3 {
		return nums[0]
	}
	return nums[2]
}
