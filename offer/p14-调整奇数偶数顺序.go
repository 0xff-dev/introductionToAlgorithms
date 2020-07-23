package offer

func ReorderOddEven(nums []int, judge func(int) bool) {
	if len(nums) == 0 {
		return
	}
	start, end := 0, len(nums)-1
	for start < end {
		for start < end && judge(nums[end]) {
			end--
		}
		for start < end && !judge(nums[start]) {
			start++
		}
		nums[start], nums[end] = nums[end], nums[start]
	}
}
