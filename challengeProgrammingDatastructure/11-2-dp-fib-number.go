package challengeprogrammingdatastructure

func Fib(n int) int {
	if n <= 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

func FibInMem(n int) int {
	if n <= 1 {
		return 1
	}
	a, b := 1, 1
	// 1, 1, 2, 3,5
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
