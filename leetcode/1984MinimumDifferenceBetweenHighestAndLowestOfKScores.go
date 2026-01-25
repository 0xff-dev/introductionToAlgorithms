package leetcode

import "sort"

func minimumDifference1984(nums []int, k int) int {
	if k == 1 {
		return 0
	}

	ret, end := 100001, len(nums)-1
	sort.Ints(nums)

	for i := 0; i <= end; i++ {
		ret = min(ret, nums[i+k-1]-nums[i])
	}
	return ret
}
