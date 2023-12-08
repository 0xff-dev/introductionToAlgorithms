package leetcode

func xorOperation(n int, start int) int {
	ans := 0
	for i := 0; i < n; i++ {
		ans ^= (start + i*2)
	}
	return ans
}
