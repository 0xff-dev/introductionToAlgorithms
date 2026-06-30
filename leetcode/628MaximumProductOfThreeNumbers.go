package leetcode

import "sort"

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	idx := sort.Search(n, func(i int) bool {
		return nums[i] >= 0
	})
	if idx == n || idx <= 1 {
		return nums[n-1] * nums[n-2] * nums[n-3]
	}
	return max(nums[0]*nums[1]*nums[n-1], nums[n-1]*nums[n-2]*nums[n-3])
}
