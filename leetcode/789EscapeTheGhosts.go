package leetcode

func distance789(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return a + b
}

func escapeGhosts(ghosts [][]int, target []int) bool {
	a, b := target[0], target[1]
	cmp := distance789(a, b)
	for _, g := range ghosts {
		a, b = target[0]-g[0], target[1]-g[1]
		tmp := distance789(a, b)
		if tmp <= cmp {
			return false
		}
	}
	return true
}
