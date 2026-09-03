package leetcode

func uniformArray3876(nums1 []int) bool {
	minOdd, minElem := 0x7fffffff, 0x7fffffff
	for i := range nums1 {
		minElem = min(minElem, nums1[i])
		if nums1[i]&1 == 0 {
			continue
		}
		minOdd = min(minOdd, nums1[i])
	}

	flag := minElem & 1
	for i := range nums1 {
		if nums1[i]&1 == flag {
			continue
		}
		if nums1[i]-minOdd >= 1 {
			continue
		}
		return false
	}
	return true
}
