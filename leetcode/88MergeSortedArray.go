package leetcode

// 1 2 3 0 0 0, 2, 5, 6
// 1, 3, 4.... 2, 2
func merge88(nums1 []int, m int, nums2 []int, n int) {
	idx := len(nums1) - 1
	i, j := m-1, n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[idx] = nums1[i]
			i--
		} else {
			nums1[idx] = nums2[j]
			j--
		}
		idx--
	}

	for ; i >= 0; i, idx = i-1, idx-1 {
		nums1[idx] = nums1[i]
	}
	for ; j >= 0; j, idx = j-1, idx-1 {
		nums1[idx] = nums2[j]
	}
}
