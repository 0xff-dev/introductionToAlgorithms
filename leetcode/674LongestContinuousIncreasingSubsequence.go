package leetcode

func findLengthOfLCIS(nums []int) int {
	ml := 1
	ans := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			ml++
			ans = max(ans, ml)
			continue
		}
		ml = 1
	}
	return ans
}
