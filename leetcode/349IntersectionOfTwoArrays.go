package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	ans := make([]int, 0)
	buckets := make([]int, 1001)
	for _, n := range nums1 {
		buckets[n] = 1
	}
	for _, n := range nums2 {
		if buckets[n] == 1 {
			ans = append(ans, n)
			buckets[n]++
		}
	}
	return ans
}
