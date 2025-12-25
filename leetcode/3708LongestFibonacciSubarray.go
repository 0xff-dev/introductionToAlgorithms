package leetcode

func longestSubarray3708(nums []int) int {
	var ret int = 2
	continued := 2
	l := len(nums)
	// 1, 2, 4, 6
	for end := 2; end < l; end++ {
		if nums[end-2]+nums[end-1] == nums[end] {
			continued++
			continue
		}
		ret = max(ret, continued)
		continued = 2
	}
	return max(ret, continued)
}
