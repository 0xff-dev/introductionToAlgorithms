package leetcode

import "sort"

func maxIceCream(costs []int, coins int) int {
	ans := 0
	sum := 0
	sort.Ints(costs)
	for i := 0; i < len(costs); i++ {
		sum += costs[i]
		if sum > coins {
			break
		}
		ans++
	}
	return ans
}
