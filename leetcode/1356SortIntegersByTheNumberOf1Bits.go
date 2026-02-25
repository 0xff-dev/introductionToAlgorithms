package leetcode

import "sort"

func bitOfOne(n int) int {
	ans := 0
	for n > 0 {
		ans++
		n = n & (n - 1)
	}
	return ans
}
func sortByBits(arr []int) []int {
	size := len(arr)
	ones := make([]int, size)
	for i := range ones {
		ones[i] = bitOfOne(arr[i])
	}
	indies := make([]int, size)
	for i := range ones {
		indies[i] = i
	}
	sort.Slice(indies, func(i, j int) bool {
		ii, jj := indies[i], indies[j]
		if ones[ii] == ones[jj] {
			return arr[ii] < arr[jj]
		}
		return ones[ii] < ones[jj]
	})
	for i := range indies {
		ones[i] = arr[indies[i]]
	}
	return ones
}
