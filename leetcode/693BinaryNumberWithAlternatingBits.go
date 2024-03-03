package leetcode

func hasAlternatingBits(n int) bool {
	cur := -1
	for n > 0 {
		oneOrZero := n & 1
		if cur == -1 || cur != oneOrZero {
			cur = oneOrZero
			n >>= 1
			continue
		}
		return false
	}
	return true
}
