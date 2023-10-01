package leetcode

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// 从数字大小来看，预计是溢出的坑
	// 如果直接n^4不出意外会直接TLE
	// 所以这里先存俩和
	l := len(nums1)
	cache := make(map[int]int)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			cache[nums1[i]+nums2[j]]++
		}
	}
	ans := 0
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			target := 0 - nums3[i] - nums4[j]
			if v, ok := cache[target]; ok {
				ans += v
			}
		}
	}
	return ans
}
