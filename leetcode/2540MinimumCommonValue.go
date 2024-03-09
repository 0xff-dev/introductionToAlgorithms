package leetcode

import "sort"

func getCommon(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			return nums1[i]
		}
		if nums1[i] < nums2[j] {
			i++
			continue
		}
		j++
	}
	return -1
}

func getCommon1(nums1 []int, nums2 []int) int {
	l2 := len(nums2)
	for _, n := range nums1 {
		if idx := sort.Search(l2, func(i int) bool {
			return nums2[i] >= n
		}); idx != l2 && nums2[idx] == n {
			return n
		}
	}
	return -1
}
