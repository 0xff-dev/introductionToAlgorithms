package leetcode

import "sort"

func chalkReplacer(chalk []int, k int) int {
	l := len(chalk)
	ans := 0
	for i := 1; i < l; i++ {
		chalk[i] += chalk[i-1]
	}
	if k > chalk[l-1] {
		k %= chalk[l-1]
	}

	for k > 0 {
		idx := sort.Search(l, func(i int) bool {
			return chalk[i] >= k
		})
		if idx == l {
			k -= chalk[l-1]
			continue
		}
		ans = idx
		if chalk[idx] == k {
			ans = (idx + 1) % l
		}
		break
	}
	return ans
}
