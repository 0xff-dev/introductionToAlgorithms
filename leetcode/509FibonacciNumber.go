package leetcode

func fib(n int) int {
	a, b := 0, 1
	if n == 0 {
		return a
	}
	if n == 1 {
		return b
	}

	// a  b
	// 0, 1, 1, 2
	for idx := 2; idx <= n; idx++ {
		t := b
		b = a + b
		a = t
	}
	return b
}
