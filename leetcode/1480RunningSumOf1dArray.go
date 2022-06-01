package leetcode

func runningSum(nums []int) []int {
	for idx := 1; idx < len(nums); idx++ {
		nums[idx] += nums[idx-1]
	}
	return nums
}
