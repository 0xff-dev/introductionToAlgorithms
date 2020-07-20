package offer

func Fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

// 相关题目
// 1. 青蛙上台阶，1次走一步或者2步
// 2. 铺砖块，砖块是2*1的大小，铺成8*2的正方形，有多少种铺法。考虑砖块的横竖问题, 横着放 f(n-2), 数着放f(n-1) 即: f(n) = f(n-1) + f(n-2)
