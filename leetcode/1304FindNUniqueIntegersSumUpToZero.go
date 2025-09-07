package leetcode

func sumZero(n int) []int {
	if n == 1 {
		return []int{0}
	}

	ans := make([]int, n)
	sum := 0
	for i := 0; i < n-1; i++ {
		ans[i] = i + 1
		sum += ans[i]
	}
	ans[n-1] = -sum
	return ans
}
