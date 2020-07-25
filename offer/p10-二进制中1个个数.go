package offer

// (n-1) & n 将n与n-1k可以把n的最右边的1变成0
func CountOfOne(n int) int {
	cnt := 0
	for n > 0 {
		cnt++
		n = (n - 1) & n
	}
	return cnt
}

// 拓展判断一个数字是不是2的整数次方 n&(n-1) == 1
// 改变m二进制多少位可以得到n， 先计算 res = m ^ n, 计算res的1的个数
