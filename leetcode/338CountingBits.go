package leetcode

func countBits(n int) []int {
	result := make([]int, n+1)
	base, pow := 0, 1

	for idx := 1; idx <= n; idx++ {
		if idx == pow {
			result[idx] = 1
			pow, base = pow*2, base+1
			continue
		}

		result[idx] = base - result[pow-1-idx]
	}
	return result
}
