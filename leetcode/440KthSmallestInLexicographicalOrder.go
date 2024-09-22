package leetcode

func findKthNumber(n int, k int) int {
	curr := 1
	k--

	for k > 0 {
		step := countSteps(n, curr, curr+1)
		if step <= k {
			curr++
			k -= step
		} else {
			curr *= 10
			k--
		}
	}

	return curr
}

func countSteps(n, prefix1, prefix2 int) int {
	steps := 0
	for prefix1 <= n {
		steps += min(n+1, prefix2) - prefix1
		prefix1 *= 10
		prefix2 *= 10
	}
	return steps
}
