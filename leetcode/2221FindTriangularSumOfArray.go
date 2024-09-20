package leetcode

func triangularSum(nums []int) int {
	l := len(nums) - 1
	for ; l > 0; l-- {
		for i := 0; i < l; i++ {
			nums[i] = (nums[i] + nums[i+1]) % 10
		}
	}
	return nums[0]
}
