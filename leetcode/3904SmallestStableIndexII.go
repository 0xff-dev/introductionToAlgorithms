package leetcode

func firstStableIndex3904(nums []int, k int) int {
	n := len(nums)
	minValues := make([]int, n)
	minValues[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		minValues[i] = min(nums[i], minValues[i+1])
	}

	ret := 0
	for i := 0; i < n; i++ {
		ret = max(ret, nums[i])
		if ret-minValues[i] <= k {
			return i
		}
	}
	return -1
}
