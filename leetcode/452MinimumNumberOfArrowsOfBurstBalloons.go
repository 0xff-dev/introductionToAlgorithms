package leetcode

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	end := points[0][1]
	ans := 0
	for i := 0; i < len(points); i++ {
		if points[i][0] > end {
			ans++
			end = points[i][1]
		}
	}

	return ans + 1
}
