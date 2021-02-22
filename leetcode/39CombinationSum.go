package leetcode

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	r := make([][]int, 0)
	combinationSumAux(candidates, target, 0, []int{}, &r)
	return r
}

func combinationSumAux(candidates []int, target, now int, path []int, r *[][]int) {
	for idx, num := range candidates {
		now += num
		if now == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			*r = append(*r, append(tmp, num))
		} else if now < target {
			combinationSumAux(candidates[idx:], target, now, append(path, num), r)
		}
		now -= num
	}
}
