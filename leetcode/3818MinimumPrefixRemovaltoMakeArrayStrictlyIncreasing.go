package leetcode

func minimumPrefixLength(nums []int) int {
	i := len(nums) - 1
	// 5 3 4
	for ; i > 0; i-- {
		if nums[i] <= nums[i-1] {
			break
		}
	}
	return i
}
