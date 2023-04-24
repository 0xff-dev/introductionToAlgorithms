package challengeprogrammingdatastructure

// 2 * 8
const mod183 = 1000000007

func modPow(x, n int) int {
	if n == 0 {
		return 1
	}
	y := modPow(x, n>>1)
	y *= y
	y %= mod183
	if n&1 == 1 {
		y *= x
		y %= mod183
	}
	return y
}

func modPowerFor(x, n int) int {
	res := 1
	for n > 0 {
		if n&1 == 1 {
			res = (res * x) % mod183
		}
		x = (x * x) % mod183
		n >>= 1
	}
	return res
}
