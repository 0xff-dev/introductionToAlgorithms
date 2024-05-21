package leetcode

import "sort"

func numSmallerByFrequency(queries []string, words []string) []int {
	ans := make([]int, len(queries))
	count := make([]int, len(words))
	for idx, word := range words {
		cur := 0
		mb := byte(' ')
		for _, b := range word {
			if mb == ' ' || mb > byte(b) {
				cur = 1
				mb = byte(b)
				continue
			}
			if mb == byte(b) {
				cur++
			}
		}
		count[idx] = cur
	}
	sort.Ints(count)
	for idx, q := range queries {
		cur := 0
		mb := byte(' ')
		for _, b := range q {
			if mb == ' ' || mb > byte(b) {
				cur = 1
				mb = byte(b)
				continue
			}
			if mb == byte(b) {
				cur++
			}
		}
		targetIndex := sort.Search(len(count), func(i int) bool {
			return count[i] > cur
		})
		ans[idx] = len(count) - targetIndex

	}
	return ans
}
