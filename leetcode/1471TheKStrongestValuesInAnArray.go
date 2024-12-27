package leetcode

import "sort"

func getStrongest1(arr []int, k int) []int {
	l := len(arr)
	sort.Ints(arr)
	median := arr[(l-1)/2]
	ai, bi := 0, l-1
	ans := make([]int, k)
	index := 0
	for ; k > 0; k, index = k-1, index+1 {
		a := arr[ai] - median
		b := arr[bi] - median
		if a < 0 {
			a = -a
		}
		if b < 0 {
			b = -b
		}
		if a > b || (a == b && arr[ai] > arr[bi]) {
			ans[index] = arr[ai]
			ai++
			continue
		}
		ans[index] = arr[bi]
		bi--
	}
	return ans
}

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
