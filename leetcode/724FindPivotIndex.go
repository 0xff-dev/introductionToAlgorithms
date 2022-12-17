package leetcode

func pivotIndex(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	sum := nums[len(nums)-1]
	left := 0
	for idx := 0; idx < len(nums); idx++ {
		right := sum - nums[idx]
		if left == right {
			return idx
		}
		left = nums[idx]
	}
	return -1
}
