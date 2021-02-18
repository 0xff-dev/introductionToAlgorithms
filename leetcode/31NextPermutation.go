package leetcode

import (
	"sort"
)

func nextPermutation(nums []int) {
	length := len(nums)
	if length == 1 || length == 0 {
		return
	}
	index := length - 1
	// 14382
	for ; index >= 1 && nums[index] <= nums[index-1]; index-- {
	}
	if index == 0 {
		reverseNums(nums)
		return
	}

	for i := length - 1; i >= index; i-- {
		if nums[i] > nums[index-1] {
			nums[i], nums[index-1] = nums[index-1], nums[i]
			break
		}
	}
	sort.Ints(nums[index:])
}

func reverseNums(nums []int) {
	for s, e := 0, len(nums)-1; s < e; s, e = s+1, e-1 {
		nums[s], nums[e] = nums[e], nums[s]
	}
}
