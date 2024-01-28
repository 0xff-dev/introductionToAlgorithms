package leetcode

/*
这道题首先想到n*n的算法，通过一个map存储到i的时候，所有的的子数组和的数量
但是会超时

自数组和前缀和，如果说prefix[i] - prefix[j] == k 那么说明[j+1....i] 是满足要求的自数组，我们只需要知道 prefix[j]有多少个即可
*/
func subarraySum(nums []int, k int) int {
	ans, sum := 0, 0
	count := make(map[int]int)
	count[0] = 1
	// 1,1,1 k= 2 预期是2
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if r := count[sum-k]; r > 0 {
			ans += r
		}
		count[sum]++
	}
	return ans
}
