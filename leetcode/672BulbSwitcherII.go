package leetcode

func op1(a uint16, n int) uint16 {
	// 0b0001111111
	return a ^ uint16((1<<n)-1)
}

func op2(a uint16, n int) uint16 {
	// 0b1010101010101010
	mask := 0xAAAA & ((1 << n) - 1)
	return a ^ uint16(mask)
}

func op3(a uint16, n int) uint16 {
	// 0b0101010101010101
	mask := 0x5555 & ((1 << n) - 1)
	return a ^ uint16(mask)

}

func op4(a uint16, n int) uint16 {
	// 0b0000100010010010
	// 1,4,7,10,13,16
	// 0b1001001001001001
	mask := 0x9999 & ((1 << n) - 1)
	return a ^ uint16(mask)
}

func flipLights(n int, presses int) int {
	// 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 1, 2
	// 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1
	// 0,    2,    2,    2,    2,    2,    2
	// 0, 3,    3,    3,    3,    3     3,
	// 0, 4,       4,       4,       4
	// 最多就是12位
	n %= 12
	if n == 0 {
		n = 12
	}
	base := uint16((1 << n) - 1)

	// 0 0 1 1 1 1 1 1
	// 0 0 6 5 4 3 2 1
	r := make(map[uint16]struct{})
	operations := []func(uint16, int) uint16{op1, op3, op2, op4}

	cache := make(map[[2]uint16]struct{})
	var dfs func(uint16, uint16)
	dfs = func(cur, left uint16) {
		if left == 0 {
			r[cur] = struct{}{}
			return
		}
		if _, ok := cache[[2]uint16{cur, left}]; ok {
			return
		}
		for _, op := range operations {
			dfs(op(cur, n), left-1)
		}
		cache[[2]uint16{cur, left}] = struct{}{}
	}
	dfs(base, uint16(presses))
	return len(r)
}
