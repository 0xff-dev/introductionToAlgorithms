package leetcode

func subtractProductAndSum(n int) int {
	p, s := 1, 0
	for n > 0 {
		mod := n % 10
		p *= mod
		s += mod
		n /= 10
	}
	return p - s
}
