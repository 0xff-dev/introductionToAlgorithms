package leetcode

import (
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	ans, tmp := 0, 1
	for idx := 1; idx < len(nums); idx++ {
		r := nums[idx] - nums[idx-1]
		if r == 1 || r == 0 {
			if r == 1 {
				tmp++
			}
			continue
		}
		if tmp > ans {
			ans = tmp
		}
		tmp = 1
	}
	if tmp > ans {
		return tmp
	}
	return ans
}
