package leetcode

func findClosest(x int, y int, z int) int {
	a := x - z
	b := y - z
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	if a == b {
		return 0
	}
	if a < b {
		return 1
	}
	return 2
}
