package leetcode

import "math"

// 1: 4, 9,
// binary search? or l + r
// 4
func judgeSquareSum(c int) bool {
	end := int(math.Sqrt(float64(c)))
	x := 0
	for x <= end {
		r := x*x + end*end
		if r == c {
			return true
		}
		if r < c {
			x++
			continue
		}
		end--
	}
	return false
}
