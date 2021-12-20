package leetcode

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	minDiff := -1
	result := make([][]int, 0)
	sort.Ints(arr)
	for idx := 0; idx < len(arr)-1; idx++ {
		abs := arr[idx+1] - arr[idx]
		if minDiff == -1 || abs == minDiff {
			minDiff = abs
			result = append(result, []int{arr[idx], arr[idx+1]})
		}
		if abs < minDiff {
			minDiff = abs
			result = result[:0]
			result = append(result, []int{arr[idx], arr[idx+1]})
		}
	}
	return result
}
