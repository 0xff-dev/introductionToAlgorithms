package leetcode

import "sort"

func checkValidCuts(n int, rectangles [][]int) bool {
	sort.Slice(rectangles, func(i, j int) bool {
		a, b := rectangles[i], rectangles[j]
		if a[1] == b[1] {
			return a[3] > b[3]
		}
		return a[1] < b[1]
	})

	start, end := rectangles[0][1], rectangles[0][3]
	cnt := 1
	for i := 1; i < len(rectangles); i++ {
		s, e := rectangles[i][1], rectangles[i][3]
		if s >= end {
			cnt++
			start, end = s, e
			end = e
			continue
		}
		start = min(start, s)
		end = max(end, e)
	}
	if cnt >= 3 {
		return true
	}

	sort.Slice(rectangles, func(i, j int) bool {
		a, b := rectangles[i], rectangles[j]
		if a[0] == b[0] {
			return a[2] > b[2]
		}
		return a[0] < b[0]
	})

	start, end = rectangles[0][0], rectangles[0][2]
	cnt = 1
	for i := 1; i < len(rectangles); i++ {
		s, e := rectangles[i][0], rectangles[i][2]
		if s >= end {
			cnt++
			end = e
			continue
		}
		start = min(start, s)
		end = max(end, e)
	}
	return cnt >= 3
}
