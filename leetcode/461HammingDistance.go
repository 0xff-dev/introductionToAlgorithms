package leetcode

func hammingDistance(x int, y int) int {
	v := x ^ y
	ans := 0
	for v > 0 {
		ans++
		v = v & (v - 1)
	}
	return ans
}
