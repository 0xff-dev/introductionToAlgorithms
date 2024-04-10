package leetcode

import "sort"

func deckRevealedIncreasing(deck []int) []int {
	// 长度是奇数还是偶数
	sort.Ints(deck)
	skip := false
	ans := make([]int, len(deck))
	a, b := 0, 0
	// 2,3,5,7,11,13,17
	// 0,1,2,3,4, 5, 6
	// 2,  3,  5,    7
	for a < len(deck) {
		if ans[b] == 0 {
			if !skip {
				ans[b] = deck[a]
				a++
			}
			skip = !skip
		}
		b = (b + 1) % len(deck)
	}
	return ans
}
