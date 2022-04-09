package leetcode

import "sort"

type fq struct {
	v, c int
}

// 大根堆，
// quick select
func topKFrequent(nums []int, k int) []int {
	numCount := make(map[int]int)
	for _, n := range nums {
		numCount[n]++
	}
	sortObj := make([]fq, 0)
	for k, v := range numCount {
		sortObj = append(sortObj, fq{k, v})
	}
	sort.Slice(sortObj, func(i, j int) bool {
		return sortObj[i].c < sortObj[j].c
	})

	result := make([]int, 0)
	for idx := len(sortObj) - 1; idx >= 0 && k > 0; idx, k = idx-1, k-1 {
		result = append(result, sortObj[idx].v)
	}
	return result
}
