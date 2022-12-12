package leetcode

import "sort"

func countElements(nums []int) int {
	ans := 0
	sort.Ints(nums)
	check := make(map[int]int)
	for _, n := range nums {
		check[n]++
	}
	if len(check) > 2 {
		return len(nums) - check[nums[0]] - check[nums[len(nums)-1]]
	}
	// 1, 2,2,2,,3,3,3,4
	// 1,,1,1,1,1,1,
	// 1,2,2,2,2
	/// ,1,1,1,2
	// 2,7,11,15

	// -3,3,3,90
	// -6 -6 -6 -5 5 5
	diff := 1
	for idx := 1; idx < len(nums); idx++ {
		if nums[idx] != nums[idx-1] {
			diff++
		}
		if diff > 1 && idx != len(nums)-1 {
			ans++
		}
	}
	if diff > 2 {
		if nums[len(nums)-1] == nums[len(nums)-2] {
			ans--
		}
		return ans
	}
	return 0
}
