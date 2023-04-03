package leetcode

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	// 哈哈哈，我傻了, 都排序了, a,b,c,d a+c <= b+d
	// b + d <= z, 可以得出a+d<=z,  且不可能出现b+c>z的情况
	ans := 0
	left, right := 0, len(people)-1
	for left <= right {
		ans++
		if people[left]+people[right] <= limit {
			left++
		}
		right--
	}

	return ans
}
