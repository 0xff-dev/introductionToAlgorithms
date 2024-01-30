package leetcode

// 1,0,1,0,1 - goal=2
func numSubarraysWithSum(nums []int, goal int) int {
	count := make(map[int]int)
	count[0] = 1
	ans := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		diff := goal - sum
		if diff > 0 {
			count[sum]++
			continue
		}
		diff = -diff
		if r := count[diff]; r > 0 {
			ans += r
		}
		count[sum]++
	}

	return ans
}
