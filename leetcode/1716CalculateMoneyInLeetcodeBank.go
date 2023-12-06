package leetcode

func totalMoney(n int) int {
	ans := 0
	a := n / 7
	aa := a + 1
	b := n % 7
	if a >= 1 {
		ans = a*28 + (a*(a-1))/2*7
	}

	for i := 0; i < b; i++ {
		ans += aa
		aa++
	}

	return ans
}
