package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	index := make(map[int]int)
	for i, n := range nums {
		if idx, ok := index[n]; ok {
			if i-idx <= k {
				return true
			}
			index[n] = i
		}
		index[n] = i
	}
	return false
}
