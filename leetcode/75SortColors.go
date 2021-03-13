package leetcode

func sortColors(nums []int) {
	length := len(nums)
	start, end := 0, length-1
	idx := 0
	for ; idx <= end; idx++ {
		if nums[idx] == 2 {
			nums[idx], nums[end] = nums[end], nums[idx]
			end--
			idx--
		} else if nums[idx] == 0 {
			nums[idx], nums[start] = nums[start], nums[idx]
			start++
		}
	}
}
