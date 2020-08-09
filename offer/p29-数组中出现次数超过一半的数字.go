package offer

func MoreThanHalf(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	result := nums[0]
	times := 1
	for index := 1; index < len(nums); index++ {
		if times == 0 {
			result = nums[index]
			times = 1
		} else if nums[index] == result {
			times++
		} else {
			times--
		}
	}
	return result
}
