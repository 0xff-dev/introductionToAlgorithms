package leetcode

// 1,2,3
// 1,2,3,4
// 2,3
// 2,3,4
// 3,4,5,6,7
// 21, 22
func countOdds(low int, high int) int {
	numCount := high - low + 1
	x := numCount / 2
	if numCount&1 == 0 {
		return x
	}
	if low&1 == 1 {
		x++
	}
	return x
}
