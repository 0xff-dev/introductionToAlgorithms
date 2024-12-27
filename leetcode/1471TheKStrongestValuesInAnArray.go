package leetcode

import "sort"

func getStrongest(arr []int, k int) []int {
	l := len(arr)
	sorted := make([]int, l)
	copy(sorted, arr)
	sort.Ints(sorted)
	median := sorted[(l-1)/2]

	sort.Slice(arr, func(i, j int) bool {
		a := arr[i] - median
		b := arr[j] - median
		if a < 0 {
			a = -a
		}
		if b < 0 {
			b = -b
		}
		if a == b {
			return arr[i] > arr[j]
		}
		return a > b
	})
	return arr[:k]
}
