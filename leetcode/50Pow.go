package leetcode

func myPow50(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	x1 := myPow50(x, n>>1)
	x1 *= x1
	if n&1 == 1 {
		x1 *= x
	}
	return x1
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	reverse := false
	if n < 0 {
		n = ^n + 1
		reverse = true
	}
	ans := myPow50(x, n)
	if reverse {
		ans = float64(1) / ans
	}
	return ans
}
