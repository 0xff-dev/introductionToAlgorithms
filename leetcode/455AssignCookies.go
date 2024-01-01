package leetcode

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	ans := 0
	for ; i < len(g) && j < len(s); j++ {
		if s[j] >= g[i] {
			i++
			ans++
		}
	}
	return ans
}
