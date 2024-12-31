package leetcode

func numEquivDominoPairs(dominoes [][]int) int {
	cache := make(map[[2]int]int)
	for _, p := range dominoes {
		a, b := p[0], p[1]
		if a > b {
			a, b = b, a
		}
		cache[[2]int{a, b}]++
	}
	ans := 0
	for _, c := range cache {
		ans += c * (c - 1) / 2
	}
	return ans
}
