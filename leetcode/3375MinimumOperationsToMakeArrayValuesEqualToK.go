package leetcode

import "sort"

func minOperations3375(nums []int, k int) int {
	used := make(map[int]struct{})
	list := make([]int, 0)
	for _, n := range nums {
		if _, ok := used[n]; !ok {
			used[n] = struct{}{}
			list = append(list, n)
		}
	}
	l := len(list)
	sort.Slice(list, func(i, j int) bool {
		return list[i] > list[j]
	})
	if list[l-1] < k {
		return -1
	}
	return sort.Search(l, func(i int) bool {
		return list[i] <= k
	})
}
