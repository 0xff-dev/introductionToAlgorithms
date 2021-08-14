package leetcode

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	cmp := make(map[int]int)
	for _, a2 := range arr2 {
		cmp[a2] = 0
	}

	leftNums := make([]int, 0)
	for _, a1 := range arr1 {
		if _, ok := cmp[a1]; ok {
			cmp[a1]++
			continue
		}

		leftNums = append(leftNums, a1)
	}

	res := make([]int, 0)
	for _, a2 := range arr2 {
		n := cmp[a2]
		for ; n > 0; n-- {
			res = append(res, a2)
		}
	}

	if len(leftNums) > 0 {
		sort.Ints(leftNums)
		res = append(res, leftNums...)
	}

	return res
}
