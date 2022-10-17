package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))

	greater := make([]int, len(nums2))
	bucket := make(map[int]int)
	length := len(nums2)
	// last don't have greater
	greater[length-1] = -1
	bucket[nums2[length-1]] = length - 1

	for idx := len(nums2) - 2; idx >= 0; idx-- {
		bucket[nums2[idx]] = idx
		if nums2[idx] < nums2[idx+1] {
			greater[idx] = idx + 1
			continue
		}
		greater[idx] = -1
		next := greater[idx+1]
		for ; next != -1 && nums2[next] < nums2[idx]; next = greater[next] {
		}
		if next != -1 {
			greater[idx] = next
		}
	}
	// 1, 2, -1, -1
	for idx, n := range nums1 {
		nextIdx := greater[bucket[n]]
		if nextIdx == -1 {
			ans[idx] = nextIdx
			continue
		}
		ans[idx] = nums2[nextIdx]
	}
	return ans
}
