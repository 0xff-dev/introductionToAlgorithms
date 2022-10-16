package leetcode

func sortArrayByParityII(nums []int) []int {
	ans := make([]int, len(nums))
	e, o := 0, 1

	// 4, 2, 5, 7
	// 0 1 2 3
	//
	for idx := 0; idx < len(nums); idx++ {
		if nums[idx]&1 == 1 {
			ans[o] = nums[idx]
			o += 2
			continue
		}

		ans[e] = nums[idx]
		e += 2
	}

	return ans
}
