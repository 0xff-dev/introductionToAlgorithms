package leetcode

import "sort"

func countSubarrays2302(nums []int, k int64) int64 {
	// n^2会超时
	// 还是前缀和
	// 一直是递增的
	ans := int64(0)
	for j := 1; j < len(nums); j++ {
		nums[j] += nums[j-1]
	}
	for j := 0; j < len(nums); j++ {
		l := j + 1
		idx := sort.Search(l, func(i int) bool {
			width := l - i
			del := 0
			if i > 0 {
				del = nums[i-1]
			}
			return int64(nums[j]-del)*int64(width) < k
		})
		if idx != l {
			ans += int64(l - idx)
		}
	}
	return ans
}
