package leetcode

func gcd914(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd914(b, a%b)
}

func hasGroupsSizeX(deck []int) bool {
	cnt := make(map[int]int)
	for i := range deck {
		cnt[deck[i]]++
	}
	cmp := 0
	for key := range cnt {
		cmp = gcd914(cmp, cnt[key])
	}
	return cmp > 1
}
