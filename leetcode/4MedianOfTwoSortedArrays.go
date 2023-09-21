package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	targetIndex := (l1 + l2) / 2
	index := 0
	i, j := 0, 0
	a, b := 0, 0
	for i < l1 && j < l2 && index <= targetIndex {
		index++
		b = a
		if nums1[i] < nums2[j] {
			a = nums1[i]
			i++
			continue
		}
		a = nums2[j]
		j++
	}
	for ; i < l1 && index <= targetIndex; i++ {
		b = a
		a = nums1[i]
		index++
	}
	for ; j < l2 && index <= targetIndex; j++ {
		b = a
		a = nums2[j]
		index++
	}
	if (l1+l2)&1 == 1 {
		return float64(a)
	}
	return float64(a+b) / 2
}
