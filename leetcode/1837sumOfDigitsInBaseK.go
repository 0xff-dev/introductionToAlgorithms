package leetcode

// 34%6=4, 5%6=5
func sumBase(n int, k int) int {

	ans := 0
	for base := n; base != 0; base /= k {
		ans += base % k
	}
	return ans
}
