package leetcode

var (
	rotateMap = map[int]int{
		0: 0, 1: 1, 8: 8,
		2: 5, 5: 2, 6: 9, 9: 6,
	}
)

func isNumOk(n int) bool {
	ans := false
	for n > 0 {
		x := n % 10
		v, ok := rotateMap[x]
		if !ok {
			// 不可以
			return false
		}
		ans = ans || x != v
		ok = ok || x != v
		n /= 10
	}
	return ans
}
func rotatedDigits(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		if isNumOk(i) {
			ans++
		}
	}
	return ans
}
