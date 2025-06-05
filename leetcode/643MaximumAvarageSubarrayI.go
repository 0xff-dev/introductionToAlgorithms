package leetcode

func findMaxAverage(nums []int, k int) float64 {
	var ans float64
	sum := 0
	s, e := 0, 0
	for ; e < k; e++ {
		sum += nums[e]
	}
	ans = float64(sum) / float64(k)
	for ; e < len(nums); s, e = s+1, e+1 {
		sum -= nums[s]
		sum += nums[e]
		ans = max(ans, float64(sum)/float64(k))
	}
	return ans
}
