package leetcode

import "sort"

// 一个人在其所有学术文章中有N篇论文分别被引用了至少N次，他的H指数就是N
func hIndex(citations []int) int {
	length := len(citations)
	// 0, 1,3,5,6 -5
	idx := sort.Search(length, func(i int) bool {
		return citations[i] >= (length - i)
	})
	return length - idx
}
