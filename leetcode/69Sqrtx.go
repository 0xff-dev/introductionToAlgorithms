package leetcode

import "sort"

func mySqrt(x int) int {
	idx := sort.Search(x, func(i int) bool {
		return i*i >= x
	})
	if idx*idx == x {
		return idx
	}
	return idx - 1
}
