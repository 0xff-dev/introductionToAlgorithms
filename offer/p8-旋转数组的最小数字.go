package offer

// 未考虑重复数据情况
func MinNumberInRotateArray(nums []int) int {
	start, end := 0, len(nums)-1
	index := start
	for nums[start] >= nums[end] {
		if end-start == 1 {
			index = end
			break
		}
		index = (start + end) / 2
		// TODO 如果num[start]==num[index] && num[index]==num[end]  直接中间寻找即可
		if nums[start] <= nums[index] {
			start = index
		} else {
			end = index
		}
	}
	return nums[index]
}
