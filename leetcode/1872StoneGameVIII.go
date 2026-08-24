package leetcode

func stoneGameVIII(stones []int) int {
	n := len(stones)
	pre := make([]int, n)
	pre[0] = stones[0]
	for i := 1; i < n; i++ {
		pre[i] = pre[i-1] + stones[i]
	}

	f := make([]int, n)
	f[n-1] = pre[n-1]
	for i := n - 2; i >= 1; i-- {
		f[i] = max(f[i+1], pre[i]-f[i+1])
	}
	return f[1]
}
