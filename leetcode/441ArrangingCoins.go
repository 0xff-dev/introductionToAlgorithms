package leetcode

import "sort"

func arrangeCoins(n int) int {
	if n == 1 {
		return 1
	}
	index := sort.Search(n, func(i int) bool {
		a := i
		b := (i + 1)
		if a&1 == 0 {
			a /= 2
		} else {
			b /= 2
		}
		return a*b > n
	})
	return index - 1
}
