package leetcode

import "sort"

type pair struct {
	idx, count int
}

func kWeakestRows(mat [][]int, k int) []int {
	rows, cols := len(mat), len(mat[0])
	// top k 问题
	pairs := make([]pair, rows)
	for r := 0; r < rows; r++ {
		sum := 0
		for c := 0; c < cols && mat[r][c] == 1; c++ {
			sum++
		}
		pairs[r] = pair{idx: r, count: sum}
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count == pairs[j].count {
			return pairs[i].idx < pairs[j].idx
		}
		return pairs[i].count < pairs[j].count
	})
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = pairs[i].idx
	}
	return ans
}
