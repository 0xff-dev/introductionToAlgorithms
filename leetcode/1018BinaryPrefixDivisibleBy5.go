package leetcode

func prefixesDivBy5(nums []int) []bool {
	base := make([]int, len(nums))
	base[0] = nums[0]
	result := make([]bool, len(nums))
	result[0] = nums[0] == 0
	for idx := 1; idx < len(nums); idx++ {
		base[idx] = nums[idx] + nums[idx-1]*2
		if idx >= 4 {
			base[idx] += base[idx-4]
		}
		cmp := base[idx]
		if idx >= 2 {
			cmp -= base[idx-2]
			if cmp < 0 {
				cmp = -cmp
			}
		}
		if cmp%5 == 0 {
			result[idx] = true
		}
	}
	return result
}
