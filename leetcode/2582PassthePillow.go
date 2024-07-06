package leetcode

func passThePillow(n int, time int) int {
	need := n - 1

	loop := time / need
	left := time % need
	if left != 0 {
		loop++
	}

	if left == 0 {
		left = need
	}

	// 1, 2, 3
	if loop&1 != 0 {
		return left + 1
	}
	return need - left + 1
}
