package leetcode

import "sort"

func buyChoco(prices []int, money int) int {
	sort.Ints(prices)
	diff := money - prices[0] - prices[1]
	if diff < 0 {
		return money
	}
	return diff
}
