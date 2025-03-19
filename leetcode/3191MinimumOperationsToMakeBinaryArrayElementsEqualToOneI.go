package leetcode

func minOperations3191(nums []int) int {
	ans := 0
	index := 0
	l := len(nums)
	// 0, 1, 0, 1, 0,
	// 1, 0, 1
	for ; index < l; index++ {
		if nums[index] == 1 {
			continue
		}
		if index >= l-2 {
			return -1
		}
		nums[index] = 1 - nums[index]
		nums[index+1] = 1 - nums[index+1]
		nums[index+2] = 1 - nums[index+2]
		ans++
	}
	return ans
}
