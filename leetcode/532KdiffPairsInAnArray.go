package leetcode

import "sort"

func findPairs(nums []int, k int) int {
	sort.Ints(nums)
	ans := 0
	for i := 0; i < len(nums)-1; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		idx := sort.Search(len(nums), func(ii int) bool {
			return ii > i && nums[ii] >= nums[i]+k
		})
		if idx != len(nums) && nums[idx] == nums[i]+k {
			ans++
		}
	}
	return ans
}
