package leetcode

func minTimeToVisitAllPoints(points [][]int) int {
	ans := 0
	for idx := 1; idx < len(points); idx++ {
		x := points[idx][0] - points[idx-1][0]
		y := points[idx][1] - points[idx-1][1]
		if x < 0 {
			x = -x
		}
		if y < 0 {
			y = -y
		}
		if x == y {
			ans += x
			continue
		}

		tmp := x
		diff := y - x
		if y < x {
			tmp = y
			diff = x - y
		}
		ans += tmp + diff
	}
	return ans
}
