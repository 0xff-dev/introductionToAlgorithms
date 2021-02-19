package leetcode

import (
	"sort"
)

// [[2,3],[4,5],[6,7],[8,9],[1,10]]
// [[1, 10]]
func merge(intervals [][]int) [][]int {
	r := make([][]int, 0)
	if len(intervals) == 0 || len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	cur := intervals[0]
	for idx := 1; idx < len(intervals); idx++ {
		if intervals[idx][1] < cur[0] {
			r = append(r, cur)
			cur = intervals[idx]
			if idx == len(intervals)-1 {
				r = append(r, cur)
			}
			continue
		}

		if intervals[idx][0] <= cur[1] {
			cur[0] = min1(min1(intervals[idx][0], cur[1]), cur[0])
			cur[1] = max1(intervals[idx][1], cur[1])
			if idx == len(intervals)-1 {
				r = append(r, cur)
			}
			continue
		}
		r = append(r, cur)
		cur = intervals[idx]
		if idx == len(intervals)-1 {
			r = append(r, cur)
		}
	}
	return r
}

func min1(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
