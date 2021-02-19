package leetcode

import (
	"sort"
)

func maxWidthOfVerticalArea(points [][]int) int {
	maxWidth := 0
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	startIndex := points[0][0]
	for idx := 1; idx < len(points); idx++ {
		if points[idx][0] == startIndex {
			continue
		}
		if points[idx][0]-startIndex > maxWidth {
			maxWidth = points[idx][0] - startIndex
		}
		startIndex = points[idx][0]
	}
	return maxWidth
}
