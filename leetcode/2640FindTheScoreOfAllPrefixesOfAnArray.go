package leetcode

func findPrefixScore(nums []int) []int64 {
	ans := make([]int64, len(nums))
	m := nums[0]
	ans[0] = int64(nums[0]) + int64(m)
	for i := 1; i < len(nums); i++ {
		m = max(m, nums[i])
		ans[i] = int64(nums[i]) + int64(m) + ans[i-1]
	}
	return ans
}
