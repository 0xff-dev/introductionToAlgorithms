package leetcode

// 有几个相同的1
func whichBase(n int) int {
	for {
		x := n & (n - 1)
		if x == 0 {
			break
		}
		n = x
	}
	return n
}

func rangeBitwiseAnd(left int, right int) int {
	// 只有在同一个范围的数字才不会出现0
	// 1, 2, 3, 4, 5, 6, 7, 8, 9
	// 0  1, 1, 2, 2, 2, 2, 3, 3

	l := whichBase(left)
	r := whichBase(right)
	if l == r {
		s := left
		for next := left + 1; next <= right; next++ {
			s &= next
		}
		return s
	}
	return 0
}
