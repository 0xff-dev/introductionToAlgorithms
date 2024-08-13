package leetcode

import (
	"fmt"
	"sort"
)

/*
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	exist := make(map[string]struct{})
	ans := make([][]int, 0)
	var dfs func(int, int, int, []int, string)
	dfs = func(index, l, sum int, path []int, key string) {
		if l == 0 {
			if sum == target {
				if _, ok := exist[key]; !ok {
					dst := make([]int, len(path))
					copy(dst, path)
					ans = append(ans, dst)
					exist[key] = struct{}{}
				}
			}
			return
		}
		if index >= len(candidates) {
			return
		}

		if sum > target {
			return
		}
		dfs(index+1, l, sum, path, key)
		dfs(index+1, l-1, sum+candidates[index], append(path, candidates[index]), key+fmt.Sprintf("-%d", candidates[index]))
	}
	for l := 1; l <= len(candidates); l++ {
		dfs(0, l, 0, []int{}, "")
	}
	return ans
}
*/

func combinationSum2(candidates []int, target int) [][]int {
	var results [][]int
	var current []int
	sort.Ints(candidates)
	findCombinations(candidates, target, 0, current, &results)
	return results
}

func findCombinations(candidates []int, target, start int, current []int, results *[][]int) {
	if target == 0 {
		temp := make([]int, len(current))
		copy(temp, current)
		*results = append(*results, temp)
		return
	}

	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		if candidates[i] > target {
			break
		}
		current = append(current, candidates[i])
		findCombinations(candidates, target-candidates[i], i+1, current, results)
		current = current[:len(current)-1]
	}
}
